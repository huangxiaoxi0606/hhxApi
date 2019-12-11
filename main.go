package main

import (
	"github.com/gin-gonic/gin"
	"hhxApi/Handlers"
	"hhxApi/Middlewares"
)

func main()  {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	//router.Use(Middlewares.CORSMiddleware())
	v1 := router.Group("/v1")
	{
		v1.GET("/aes",Middlewares.SetUp(), Handlers.AesTest)

	}
	router.Run(":8088")
}

//func Entry() gin.HandlerFunc {
//	return func(context *gin.Context) {
//		utilGin := Until.Gin{Ctx:context}
//		appKey, ok := context.GetQuery("app_key")
//		if !ok {
//			utilGin.Response(Config.Code_Param_Must, "app_key is required", nil)
//			return
//		}
//		appSecret, ok := context.GetQuery("app_secret")
//		if !ok {
//			context.String(203, "app_secret is required")
//			return
//		}
//		encryptStr := "ak=" + appKey + "&ts=1111111111"
//		// 生成签名
//		sn, _ := Helpers.AesEncrypt(encryptStr, []byte(appSecret), appSecret)
//		// 验证签名
//		Helpers.AesDecrypt(sn, []byte(appSecret), appSecret)
//		context.Next()
//	}
//}


