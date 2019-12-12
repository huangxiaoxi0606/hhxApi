/*
@Time : 2019/12/12 9:28
@Author : hhx06
@File : DirectionLogHandler
@Software: GoLand
*/
package Handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"hhxApi/Config"
	"hhxApi/Helpers"
	"hhxApi/Response"
	"hhxApi/Until"
	"log"
)

func GetDirectionWeekLog(context *gin.Context) {
	db, err := gorm.Open("mysql", Config.DSN)
	defer db.Close()
	if err != nil {
		log.Panic("mysql db connect faild --- " + err.Error())
	}
	var monday = Helpers.GetFirstDateOfWeek()
	var result = []Response.DirectionLogResponce{}
	db.Table("direction_logs").Select("direction_logs.illustration,direction_logs.id," +
		"direction_logs.daily_id,direction_logs.status,direction_logs.ok,direction_logs.money,direction_logs.week_day," +
		"direction_logs.week_day,direction_logs.created_at, directions.name").Joins("left join directions on directions.id = direction_logs.direction_id").Where(" direction_logs.created_at > ?", monday).Scan(&result)
	utilGin := Until.Gin{Ctx:context}
	utilGin.Response(Config.STATUS_OK, "获取成功", result)
}