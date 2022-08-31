package main

import (
	_ "crypto/hmac"
	_ "crypto/sha1"
	_ "encoding/base64"
	"fmt"
	"io/ioutil"
	"math/rand"
	_ "net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/TTvcloud/vcloud-sdk-golang/base"
	"github.com/TTvcloud/vcloud-sdk-golang/service/imagex"
)

/*
 * upload local images
 */

var (
	envFlag      = 1 // 1 表示在boe环境测试， 0表示在线上测试
	domainExp, _ = regexp.Compile("[a-zA-z]+://[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(\\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+\\.?")
)

func main() {
	// default region cn-north-1, for other region, call imagex.NewInstanceWithRegion(region)

	instance1 := getImgXInstance(envFlag)
	var serviceId string
	var mainDomain string
	var proto string
	prodOpt := []imagex.OptionFun{imagex.WithHttps(), imagex.WithFormat(imagex.FORMAT_PNG)}
	boeOpt := []imagex.OptionFun{imagex.WithFormat(imagex.FORMAT_PNG)}
	var opt []imagex.OptionFun
	//项目已上线，拒绝在线上测试
	if envFlag == 0 && false {
		serviceId = "zqfot2bwd8"
		r := rand.New(rand.NewSource(time.Now().Unix()))
		numStr := strconv.FormatInt(int64((r.Intn(3)+1)*3), 10)
		mainDomain = "p" + numStr + "-vxe.byteimg.com"
		proto = "http"
		opt = prodOpt
	} else {
		serviceId = "vf9i25tua1"
		mainDomain = "p-boe.byted.org"
		proto = "http"
		opt = boeOpt
	}

	params := &imagex.ApplyUploadImageParam{
		ServiceId: serviceId,
	}
	filePath := "/Users/bytedance/Downloads/img_test/1628151380941.png"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("read file from %s error %v", filePath, err)
		fmt.Printf("test mainDomain:%s proto:%s", mainDomain, proto)
		os.Exit(-1)
	}

	resp, err := instance1.UploadImages(params, [][]byte{data})
	if err != nil {
		fmt.Printf("error %v\n", err)
	} else {
		fmt.Printf("success %+v\n", resp)
	}

	fmt.Printf("%+v\n", opt)
	fmt.Printf("%s\n", resp.Results[0].Uri)
	uri := resp.Results[0].Uri
	tpl := "tplv-" + serviceId + "-image"

	urls, err := instance1.GetImagexURL(serviceId, uri, tpl, opt...)
	if err != nil {
		fmt.Printf("error %v\n", err)
	} else {
		fmt.Printf("MainUrl:%s uri:%s\n", urls.MainUrl, GetURIByURL(urls.MainUrl))
		fmt.Printf("BackupUrl:%s uri:%s\n", urls.BackupUrl, GetURIByURL(urls.BackupUrl))
	}

	instance2 := getImgXInstance(envFlag)
	delImg, err := instance2.DeleteImages(serviceId, []string{uri})
	if err != nil {
		fmt.Printf("error:%+v\n", err)
	} else {
		fmt.Printf("%+v", delImg.DeletedFiles)
	}
}

func getImgXInstance(envFlag int) *imagex.ImageXClient {
	instance := imagex.NewInstance()
	if envFlag == 0 && false {
		// call below method if you dont set ak and sk in ～/.vcloud/config
		//instance.SetCredential(base.Credentials{
		//	AccessKeyID:     "AKLTY2RlMDJhMzExNWQwNGM1MTk2NjNkYjI3ODE4MDhjMjY",
		//	SecretAccessKey: "ABjMCfWkSFBoiwh1Cb3Eh/QCW8Psow+fb6DkcYwBXK+hawcCro9otzR+hyKQxVsi",
		//})
	} else {

		instance.SetCredential(base.Credentials{
			AccessKeyID:     "AKLTYjM0ZTc4MTI2M2IyNDRhYzlhOWZjNjExYzYyMDEyMDk",
			SecretAccessKey: "WVRnNE5EVTNOVFZpTURZNU5EazVaams0TmpRNU1qa3hZakE0TnpVd00yRQ==",
		})
		/*instance1.SetCredential(base.Credentials{
			AccessKeyID:     "boeorigin",
			SecretAccessKey: "07ed6295f377429b9d18b72dec91f4d1",
		})*/
		instance.SetHost("staging-openapi-boe.byted.org")
	}
	return instance
}

//通过url获取uri
func GetURIByURL(url string) string {
	domain := domainExp.FindString(url)
	if domain == "" {
		return ""
	} else {
		return strings.Replace(url, domain, "", -1)
	}
}

/*{

path := fmt.Sprintf("/%s~%s.%s", uri, tpl,imagex.FORMAT_PNG)
sigTxt := path
sigTxt = path + "?"
path = sigTxt
urls2 := &imagex.ImgUrl{
	MainUrl:   fmt.Sprintf("%s://%s%s", proto, mainDomain, path),
}
fmt.Printf("%s\n", urls2.MainUrl)
fmt.Printf("%s\n", urls2.BackupUrl)
//delImg, err := instance1.DeleteImages(serviceId, []string{"abc"})
	if env.IsPPE() || env.IsProduct() {
		opt = prodOpt
		fmt.Println("ppe")
	} else {
		opt = boeOpt
		fmt.Println("boe")
	}
}*/
