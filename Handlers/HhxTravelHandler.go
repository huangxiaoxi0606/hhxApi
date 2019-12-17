/*
@Time : 2019/12/16 9:30
@Author : hhx06
@File : HhxTravelHandler
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

func HhTravelList(context *gin.Context) {
	PageNum := 0
	PageSize := 10
	maps := make(map[string]interface{})
	list, err := Models.GetTravelLists(PageNum, PageSize, maps)
	if err != nil {
		log.Panic("hhxTravelList is wrong --- " + err.Error())
		return
	}
	utilGin := Until.Gin{Ctx: context}
	utilGin.Response(http.StatusOK, "获取成功", list)

}

func HhTravelIntro(context *gin.Context) {
	utilGin := Until.Gin{Ctx: context}
	id, ok := context.GetQuery("id")
	if !ok {
		utilGin.Response(http.StatusBadRequest, "id is required", nil)
		return
	}
	ids, _ := strconv.Atoi(id)
	intro, err := Models.GetHhxTravelDetail(ids)
	if err != nil {
		log.Panic("hhxDetail is wrong--- " + err.Error())
		return
	}
	utilGin.Response(http.StatusOK, "获取成功", intro)
}

func HhTravelTraffic(context *gin.Context) {
	utilGin := Until.Gin{Ctx: context}
	id, ok := context.GetQuery("id")
	if !ok {
		utilGin.Response(http.StatusBadRequest, "id is required", nil)
		return
	}
	ids, _ := strconv.Atoi(id)
	traffics, err := Models.GetHhxTravelTraffic(ids)
	if err != nil {
		log.Panic("hhxTraffic is wrong --- " + err.Error())
		return
	}
	utilGin.Response(http.StatusOK, "获取成功", traffics)

}

func HhTravelBill(context *gin.Context) {
	utilGin := Until.Gin{Ctx: context}
	id, ok := context.GetQuery("id")
	if !ok {
		utilGin.Response(http.StatusBadRequest, "id is required", nil)
		return
	}
	ids, _ := strconv.Atoi(id)
	bills, err := Models.GetHhxTravelBill(ids)
	if err != nil {
		log.Panic("bills is wrong --- " + err.Error())
		return
	}
	utilGin.Response(http.StatusOK, "获取成功", bills)

}
