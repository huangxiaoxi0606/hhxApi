/*
@Time : 2019/12/11 9:30
@Author : hhx06
@File : ENCRYPT
@Software: GoLand
*/
package Middlewares

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-api-master/app/config"
	"hhxApi/Config"
	"hhxApi/Helpers"
	"hhxApi/Until"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

var AppSecret string

// AES 对称加密
func SetUp() gin.HandlerFunc {

	return func(c *gin.Context) {
		utilGin := Until.Gin{Ctx: c}

		sign, err := verifySign(c)

		if sign != nil {
			utilGin.Response(-1, "Wrong Sign", sign)
			c.Abort()
			return
		}

		if err != nil {
			utilGin.Response(-1, err.Error(), sign)
			c.Abort()
			return
		}

		c.Next()
	}
}

// 验证签名
func verifySign(c *gin.Context) (map[string]string, error) {
	_ = c.Request.ParseForm()
	req := c.Request.Form
	debug := strings.Join(c.Request.Form["debug"], "")
	ak := strings.Join(c.Request.Form["ak"], "")
	sn := strings.Join(c.Request.Form["sn"], "")
	ts := strings.Join(c.Request.Form["ts"], "")

	//验证来源
	value := Config.AppAk
	if ak == value {
		AppSecret = Config.AppAes
	}else{
		return nil, errors.New("ak Error")
	}

	if debug == "1" {
		currentUnix := Helpers.GetCurrentUnix()
		req.Set("ts", strconv.FormatInt(currentUnix, 10))

		sn, err := createSign(req)
		if err != nil {
			return nil, errors.New("sn Exception")
		}

		res := map[string]string{
			"ts": strconv.FormatInt(currentUnix, 10),
			"sn": sn,
		}
		return res, nil
	}

	// 验证过期时间
	timestamp := time.Now().Unix()
	exp, _ := strconv.ParseInt(config.AppSignExpiry, 10, 64)
	tsInt, _ := strconv.ParseInt(ts, 10, 64)
	if tsInt > timestamp || timestamp-tsInt >= exp {
		return nil, errors.New("ts Error")
	}

	// 验证签名
	if sn == "" {
		return nil, errors.New("sn Error")
	}

	decryptStr, decryptErr := Helpers.AesDecrypt(sn, []byte(AppSecret), AppSecret)
	if decryptErr != nil {
		return nil, errors.New(decryptErr.Error())
	}
	if decryptStr != createEncryptStr(req) {
		return nil, errors.New("sn Error")
	}
	return nil, nil
}

// 创建签名
func createSign(params url.Values) (string, error) {
	return Helpers.AesEncrypt(createEncryptStr(params), []byte(AppSecret), AppSecret)
}

func createEncryptStr(params url.Values) string {
	var key []string
	var str = ""
	for k := range params {
		if k != "sn" && k != "debug" {
			key = append(key, k)
		}
	}
	sort.Strings(key)
	for i := 0; i < len(key); i++ {
		if i == 0 {
			str = fmt.Sprintf("%v=%v", key[i], params.Get(key[i]))
		} else {
			str = str + fmt.Sprintf("&%v=%v", key[i], params.Get(key[i]))
		}
	}
	return str
}
