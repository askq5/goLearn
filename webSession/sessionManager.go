package webSession

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)




//session管理器
type SessionManager struct {
	cookieName    string          //cookie名称
	cookieExpire  int             //cookie有效期时间（单位：分钟，0表示会话cookie）
	sessionExpire int64           //session有效期时间（单位：分钟）
	gcDuration    int             //垃圾回收机制运行间隔时间（单位：分钟）
	provider      SessionProvider //session存储器
}
//创建session管理器
func NewManager(cookieName string, cookieExpire int, sessionExpire int64, gcDuration int, provider SessionProvider) *SessionManager {
	return &SessionManager{
		cookieName:    cookieName,
		cookieExpire:  cookieExpire,
		sessionExpire: sessionExpire,
		gcDuration:    gcDuration,
		provider:      provider,
	}
}



//生成session ID
func (sm *SessionManager) createSessionId(req *http.Request) string {
	addr := req.RemoteAddr
	userAgent := req.Header.Get("User-Agent")
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(10000)
	str := addr + "_" + userAgent + "_" + strconv.Itoa(n)
	h := md5.New()
	h.Write([]byte(str))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

//读cookie，获取session ID
func (sm *SessionManager) getSessionId(req *http.Request) (string, error) {
	c, err := req.Cookie(sm.cookieName)
	if err != nil {
		return "", errors.New("Reading cookie failed: " + err.Error())
	}
	if len(c.Value) == 0 { //尚未设置cookie
		return "", errors.New("Cookie does not exists: " + sm.cookieName)
	}
	return c.Value, nil
}

//创建session回话
func (sm *SessionManager) Create(writer *http.ResponseWriter, req *http.Request, data map[string]interface{}) error {
	sessionId, _ := sm.getSessionId(req)
	if len(sessionId) > 0 {
		data, _ := sm.provider.getAll(sessionId)
		if data != nil { //已有session，无需创建
			return nil
		}
	}
	sessionId = sm.createSessionId(req)
	if len(sessionId) == 0 {
		return errors.New("Length of sessionId is 0")
	}
	err := sm.provider.create(sessionId, data)
	if err != nil {
		return err
	}
	if sm.cookieExpire == 0 { //会话cookie
		http.SetCookie(*writer, &http.Cookie{
			Name:     sm.cookieName,
			Value:    sessionId,
			Path:     "/", //一定要设置为根目录，才能在所有页面生效
			HttpOnly: true,
		})
	} else { //持久cookie
		expire, _ := time.ParseDuration(strconv.Itoa(sm.cookieExpire) + "m")
		http.SetCookie(*writer, &http.Cookie{
			Name:     sm.cookieName,
			Value:    sessionId,
			Path:     "/", //一定要设置为根目录，才能在所有页面生效
			Expires:  time.Now().Add(expire),
			HttpOnly: true,
		})
	}
	return nil
}

//获取session键值
func (sm *SessionManager) Get(writer *http.ResponseWriter, req *http.Request, key string) (string, error) {
	sessionId, _ := sm.getSessionId(req)
	if len(sessionId) == 0 {
		return "", errors.New("Length of sessionId is 0")
	}
	return sm.provider.get(sessionId, key)
}

//读取session所有键值对
func (sm *SessionManager) GetAll(writer *http.ResponseWriter, req *http.Request) (map[string]string, error) {
	sessionId, _ := sm.getSessionId(req)
	if len(sessionId) == 0 {
		return nil, errors.New("Length of sessionId is 0")
	}
	return sm.provider.getAll(sessionId)
}

//设置session键值
func (sm *SessionManager) Set(writer *http.ResponseWriter, req *http.Request, key string, value interface{}) error {
	sessionId, _ := sm.getSessionId(req)
	if len(sessionId) == 0 {
		return errors.New("Length of sessionId is 0")
	}
	return sm.provider.set(sessionId, key, value)
}

//销毁session
func (sm *SessionManager) Destroy(req *http.Request) error {
	sessionId, _ := sm.getSessionId(req)
	if len(sessionId) == 0 {
		return errors.New("Length of sessionId is 0")
	}
	return sm.provider.destroy(sessionId)
}

//垃圾回收：删除过期session
func (sm *SessionManager) Gc() error {
	err := sm.provider.gc(sm.sessionExpire)
	duration, _ := time.ParseDuration(strconv.Itoa(sm.gcDuration) + "m")
	time.AfterFunc(duration, func() { sm.Gc() }) //设置下次运行时间
	return err
}

