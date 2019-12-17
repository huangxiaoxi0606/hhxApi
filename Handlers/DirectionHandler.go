/*
@Time : 2019/12/13 14:49
@Author : hhx06
@File : DirectionHandler
@Software: GoLand
*/
package Handlers

import (
	"github.com/gin-gonic/gin"
	"hhxApi/Models"
	"hhxApi/Until"
	"log"
	"net/http"
)

func DirectionList(context *gin.Context) {
	PageNum := 1
	PageSize := 10
	maps := make(map[string]interface{})
	directionList, err := Models.GetDirectionLists(PageNum, PageSize, maps)
	if err != nil {
		log.Panic("dbTopList is wrong --- " + err.Error())
		return
	}
	utilGin := Until.Gin{Ctx: context}
	utilGin.Response(http.StatusOK, "获取成功", directionList)
}
