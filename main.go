package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"hhxApi/Config"
	"hhxApi/Handlers"
	"hhxApi/Middlewares"
	"hhxApi/Models"
)

func init() {
	Models.Su()
}

func main() {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(Middlewares.CORSMiddleware())
	v1 := router.Group("/v1")
	{
		if Config.Mode == "product" {
			v1.Use(Middlewares.SetUp())
		}
		v1.GET("/aes", Handlers.AesTest)
		v1.GET("/db_top_list", Handlers.DbTopList)
		v1.GET("/db_top_detail", Handlers.DbTopDetail)
		v1.GET("/db_music_top_list", Handlers.DbMusicTopList)
		v1.GET("/db_music_top_detail", Handlers.DbMusicTopDetail)

		v1.GET("/direction_week_log", Handlers.DirectionWeekLog)
		v1.GET("/direction_today_log", Handlers.DirectionTodayLog)
		v1.GET("/direction_mouth_log", Handlers.DirectionMouthLog)

		v1.GET("/direction_list", Handlers.DirectionList)

		v1.GET("/direction_mouth_sum", Handlers.DirectionMouthSum)
		v1.GET("/direction_week_sum", Handlers.DirectionWeekSum)
		v1.GET("/directionlog_mold_id", Handlers.DirectionLogByMoldId)

		v1.GET("/hhx_travil_list", Handlers.HhxTravilList)
		v1.GET("/hhx_travil_detail", Handlers.HhxTravilDetail)
		v1.GET("/hhx_travil_traffic", Handlers.HhxTravilTraffic)
		v1.GET("/hhx_travil_bill", Handlers.HhxTravilBill)


		v1.GET("/add_have", Handlers.AddHave)

	}

	var db *gorm.DB
	defer db.Close()
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
