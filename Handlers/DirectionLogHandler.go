/*
@Time : 2019/12/12 9:28
@Author : hhx06
@File : DirectionLogHandler
@Software: GoLand
*/
package Handlers

import (
	"github.com/gin-gonic/gin"
	"hhxApi/Config"
	"hhxApi/Models"
	"hhxApi/Until"
	"log"
)

func GetDirectionWeekLog(context *gin.Context) {
	result, err1 := Models.GetLog(2)
	if err1 != nil {
		log.Panic("directionlog faild --- " + err1.Error())
	}
	utilGin := Until.Gin{Ctx: context}
	utilGin.Response(Config.STATUS_OK, "获取成功", result)
}

func GetDirectionTodayLog(context *gin.Context) {
	result, err1 := Models.GetLog(1)
	if err1 != nil {
		log.Panic("directionlog faild --- " + err1.Error())
	}
	utilGin := Until.Gin{Ctx: context}
	utilGin.Response(Config.STATUS_OK, "获取成功", result)
}


func GetDirectionMouthLog(context *gin.Context) {
	result, err1 := Models.GetLog(3)
	if err1 != nil {
		log.Panic("directionlog faild --- " + err1.Error())
	}
	utilGin := Until.Gin{Ctx: context}
	utilGin.Response(Config.STATUS_OK, "获取成功", result)
}

func GetDirectionMouthSum(context *gin.Context)  {
	result, err1 := Models.GetDirectionSum(3)
	if err1 != nil {
		log.Panic("directionlog faild --- " + err1.Error())
	}
	utilGin := Until.Gin{Ctx: context}
	utilGin.Response(Config.STATUS_OK, "获取成功", result)
}

func GetDirectionWeekSum(context *gin.Context)  {
	result, err1 := Models.GetDirectionSum(2)
	if err1 != nil {
		log.Panic("directionlog faild --- " + err1.Error())
	}
	utilGin := Until.Gin{Ctx: context}
	utilGin.Response(Config.STATUS_OK, "获取成功", result)
}

