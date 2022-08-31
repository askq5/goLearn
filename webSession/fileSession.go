package webSession

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
)

//文件session存储器
type FileProvider struct {
	savePath string                   //session文件保存路径
	muxMap   map[string]*sync.RWMutex //session文件锁
}

//创建文件session存储器对象
func NewFileProvider(savePath string) *FileProvider {
	return &FileProvider{
		savePath: savePath,
		muxMap:   make(map[string]*sync.RWMutex),
	}
}

//返回session文件名称
func (fp FileProvider) filename(sessionId string) string {
	return fp.savePath + "/" + sessionId
}
//将数据类型转换为字符串
func (fp FileProvider) toString(value interface{}) (string, error) {
	var str string
	vType := reflect.TypeOf(value)
	switch vType.Name() {
	case "int":
		i, _ := value.(int)
		str = strconv.Itoa(i)
	case "string":
		str, _ = value.(string)
	case "int64":
		i, _ := value.(int64)
		str = strconv.FormatInt(i, 10)
	default:
		return "", errors.New("Unsupported type: " + vType.Name())
	}
	return str, nil
}
//创建/重写session文件
func (fp FileProvider) write(sessionId string, data map[string]string, newFile bool) error {
	_, exist := fp.muxMap[sessionId]
	if !exist { //内存中没有锁，先建锁
		fp.muxMap[sessionId] = new(sync.RWMutex)
	}
	fp.muxMap[sessionId].Lock()
	defer func() {
		fp.muxMap[sessionId].Unlock()
	}()
	fname := fp.filename(sessionId)
	_, err := os.Stat(fname)
	var f *os.File
	if newFile {
		if err == nil { //若session文件存在，则先删除
			os.Remove(fname)
		}
		f, err = os.Create(fname)
		if err != nil {
			return errors.New("Creating session file failed: " + err.Error())
		}
	} else {
		if err != nil { //session文件不存在
			return errors.New("Session file does not exists: " + fname)
		}
		f, err = os.OpenFile(fname, os.O_RDWR|os.O_TRUNC, 0644)
		if err != nil {
			return errors.New("Opening session file failed: " + err.Error())
		}
	}
	defer func() {
		os.Chtimes(fname, time.Now(), time.Now()) //更新文件最后访问时间
		f.Close()
	}()
	for key, value := range data {
		_, err = fmt.Fprintln(f, key+":"+value)
		if err != nil {
			return errors.New("Setting session key value failed: " + err.Error())
		}
	}
	return nil
}

//读取session文件
func (fp FileProvider) read(sessionId string) (map[string]string, error) {
	fname := fp.filename(sessionId)
	_, err := os.Stat(fname)
	if err != nil { //session文件不存在
		return nil, errors.New("Session file does not exists: " + fname)
	}
	_, exist := fp.muxMap[sessionId]
	if !exist { //内存中没有锁，先建锁
		fp.muxMap[sessionId] = new(sync.RWMutex)
	}
	fp.muxMap[sessionId].Lock()
	defer func() {
		fp.muxMap[sessionId].Unlock()
	}()
	f, err := os.Open(fname)
	if err != nil {
		return nil, errors.New("Opening session file failed: " + err.Error())
	}
	defer func() {
		os.Chtimes(fname, time.Now(), time.Now()) //更新文件最后访问时间
		f.Close()
	}()
	data := make(map[string]string)
	scaner := bufio.NewScanner(f)
	for scaner.Scan() {
		kv := strings.Split(scaner.Text(), ":")
		if len(kv) != 2 {
			continue
		}
		data[kv[0]] = kv[1]
	}
	if len(data) == 0 {
		return nil, errors.New("No data in session file")
	}
	return data, nil
}

//创建session存储空间
func (fp FileProvider) create(sessionId string, data map[string]interface{}) error {
	strData := make(map[string]string)
	for key, value := range data {
		strValue, err := fp.toString(value)
		if err != nil {
			return err
		}
		strData[key] = strValue
	}
	return fp.write(sessionId, strData, true)
}

//读取session键值
func (fp FileProvider) get(sessionId, key string) (string, error) {
	data, err := fp.read(sessionId)
	if err != nil {
		return "", err
	}
	value, ok := data[key]
	if !ok {
		return "", errors.New("Session key does not exists: " + key)
	}
	return value, nil
}

//读取session所有键值对
func (fp FileProvider) getAll(sessionId string) (map[string]string, error) {
	return fp.read(sessionId)
}

//设置session键值
func (fp FileProvider) set(sessionId, key string, value interface{}) error {
	data, err := fp.read(sessionId)
	if data == nil {
		return err
	}
	str, err := fp.toString(value)
	if err != nil {
		return err
	}
	data[key] = str
	return fp.write(sessionId, data, false)
}

//销毁session：删除session文件
func (fp FileProvider) destroy(sessionId string) error {
	fname := fp.filename(sessionId)
	_, err := os.Stat(fname)
	if err != nil { //session文件不存在
		return errors.New("Session file does not exists: " + fname)
	}
	_, exist := fp.muxMap[sessionId]
	if !exist { //内存中没有锁，先建锁
		fp.muxMap[sessionId] = new(sync.RWMutex)
	}
	fp.muxMap[sessionId].Lock()
	err = os.Remove(fname)
	fp.muxMap[sessionId].Unlock()
	if err != nil {
		return errors.New("Removing session file failed: " + err.Error())
	}
	delete(fp.muxMap, sessionId)
	return nil
}

//垃圾回收：删除过期session文件
func (fp FileProvider) gc(expire int64) error {
	now := time.Now().Unix()
	for sessionId, mux := range fp.muxMap {
		fname := fp.filename(sessionId)
		if len(fname) == 0 {
			continue
		}
		mux.Lock()
		info, err := os.Stat(fname)
		if err != nil {
			mux.Unlock()
			continue
		}
		modTime := info.ModTime().Unix() //文件最后访问时间
		if modTime+expire*60 < now {     //已超出过期时间
			err = os.Remove(fname)
			mux.Unlock()
			if err != nil {
				delete(fp.muxMap, sessionId)
			}
		} else {
			mux.Unlock()
		}
	}
	return nil
}