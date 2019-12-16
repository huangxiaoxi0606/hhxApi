/*
@Time : 2019/12/16 9:30
@Author : hhx06
@File : HhxTravilHandler
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

func HhxTravilList(context *gin.Context) {
	PageNum := 0
	PageSize := 10
	maps := make(map[string]interface{})
	hhxTravilList, err := Models.GetTravilLists(PageNum, PageSize, maps)
	if err != nil {
		log.Panic("dbTopList faild --- " + err.Error())
		return
	}
	utilGin := Until.Gin{Ctx: context}
	utilGin.Response(http.StatusOK, "获取成功", hhxTravilList)

}

func HhxTravilDetail(context *gin.Context) {
	utilGin := Until.Gin{Ctx: context}
	id, ok := context.GetQuery("id")
	if !ok {
		utilGin.Response(http.StatusBadRequest, "id is required", nil)
		return
	}
	ids, _ := strconv.Atoi(id)
	hhxTravilDetail, err := Models.GetHhxTravilDetail(ids)
	if err != nil {
		log.Panic("hhxTravilDetail faild --- " + err.Error())
		return
	}
	utilGin.Response(http.StatusOK, "获取成功", hhxTravilDetail)
}

func HhxTravilTraffic(context *gin.Context)  {
	utilGin := Until.Gin{Ctx: context}
	id, ok := context.GetQuery("id")
	if !ok {
		utilGin.Response(http.StatusBadRequest, "id is required", nil)
		return
	}
	ids, _ := strconv.Atoi(id)
	hhxTravilTraffic, err := Models.GetHhxTravilTraffic(ids)
	if err != nil {
		log.Panic("hhxTravilTraffic faild --- " + err.Error())
		return
	}
	utilGin.Response(http.StatusOK, "获取成功", hhxTravilTraffic)

}


func HhxTravilBill(context *gin.Context)  {
	utilGin := Until.Gin{Ctx: context}
	id, ok := context.GetQuery("id")
	if !ok {
		utilGin.Response(http.StatusBadRequest, "id is required", nil)
		return
	}
	ids, _ := strconv.Atoi(id)
	hhxTravilBill, err := Models.GetHhxTravilBill(ids)
	if err != nil {
		log.Panic("hhxTravilTraffic faild --- " + err.Error())
		return
	}
	utilGin.Response(http.StatusOK, "获取成功", hhxTravilBill)

}

