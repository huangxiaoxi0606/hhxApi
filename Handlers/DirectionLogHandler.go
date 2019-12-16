/*
@Time : 2019/12/12 9:28
@Author : hhx06
@File : DirectionLogHandler
@Software: GoLand
*/
package Handlers

import (
	"github.com/gin-gonic/gin"
	"hhxApi/Models"
	"hhxApi/Until"
	"log"
	"net/http"
	"strconv"
)

func DirectionWeekLog(context *gin.Context) {
	result, err1 := Models.GetLog(2)
	if err1 != nil {
		log.Panic("directionlog faild --- " + err1.Error())
	}
	utilGin := Until.Gin{Ctx: context}
	utilGin.Response(http.StatusOK, "获取成功", result)
}

func DirectionTodayLog(context *gin.Context) {
	result, err1 := Models.GetLog(1)
	if err1 != nil {
		log.Panic("directionlog faild --- " + err1.Error())
	}
	utilGin := Until.Gin{Ctx: context}
	utilGin.Response(http.StatusOK, "获取成功", result)
}


func DirectionMouthLog(context *gin.Context) {
	result, err1 := Models.GetLog(3)
	if err1 != nil {
		log.Panic("directionlog faild --- " + err1.Error())
	}
	utilGin := Until.Gin{Ctx: context}
	utilGin.Response(http.StatusOK, "获取成功", result)
}

func DirectionMouthSum(context *gin.Context)  {
	result, err1 := Models.GetDirectionSum(3)
	if err1 != nil {
		log.Panic("directionlog faild --- " + err1.Error())
	}
	utilGin := Until.Gin{Ctx: context}
	utilGin.Response(http.StatusOK, "获取成功", result)
}

func DirectionWeekSum(context *gin.Context)  {
	result, err1 := Models.GetDirectionSum(2)
	if err1 != nil {
		log.Panic("directionlog faild --- " + err1.Error())
	}
	utilGin := Until.Gin{Ctx: context}
	utilGin.Response(http.StatusOK, "获取成功", result)
}

func DirectionLogByMoldId(context *gin.Context)  {
	utilGin := Until.Gin{Ctx: context}
	id, ok := context.GetQuery("id")
	if !ok {
		utilGin.Response(http.StatusBadRequest, "id is required", nil)
		return
	}
	ids, _ := strconv.Atoi(id)
	mold, ok := context.GetQuery("mold")
	if !ok {
		utilGin.Response(http.StatusBadRequest, "mold is required", nil)
		return
	}
	molds, _ := strconv.Atoi(mold)
	result, err1 := Models.GetDirectionLog(molds,ids)
	if err1 != nil {
		log.Panic("DirectionLogById faild --- " + err1.Error())
	}
	utilGin.Response(http.StatusOK, "获取成功", result)
}

