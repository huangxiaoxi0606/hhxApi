/*
@Time : 2019/12/11 14:30
@Author : hhx06
@File : DbTopHandler
@Software: GoLand
*/
package Handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"hhxApi/Models"
	"hhxApi/Until"
	"log"
	"net/http"
	"strconv"
)

func DbTopList(context *gin.Context) {
	//db, err := gorm.Open("mysql", Config.DSN)
	//defer db.Close()
	//if err != nil {
	//	log.Panic("mysql db connect faild --- " + err.Error())
	//}
	//var dbTopList = []Response.DbTopResponse{}
	//db.Where("id >?", 0).Find(&dbTopList).Scan(&dbTopList)
	//utilGin := Until.Gin{Ctx: context}
	//utilGin.Response(Config.STATUS_OK, "获取成功", dbTopList)

	PageNum := 1
	PageSize := 10
	maps := make(map[string]interface{})
	dbTopList, err := Models.GetTopLists(PageNum, PageSize, maps)
	if err != nil {
		log.Panic("dbTopList faild --- " + err.Error())
		return
	}
	utilGin := Until.Gin{Ctx: context}
	utilGin.Response(http.StatusOK, "获取成功", dbTopList)

}

func DbTopDetail(context *gin.Context) {
	utilGin := Until.Gin{Ctx: context}
	id, ok := context.GetQuery("id")
	if !ok {
		utilGin.Response(http.StatusBadRequest, "id is required", nil)
		return
	}
	ids, _ := strconv.Atoi(id)
	//var dbDeatil = Response.DbTopResponse{}
	//db, err := gorm.Open("mysql", Config.DSN)
	//defer db.Close()
	//if err != nil {
	//	log.Panic("mysql db connect faild --- " + err.Error())
	//}
	//db.Where("id =?", id).Find(&dbDeatil).Scan(&dbDeatil)
	dbDeatil, err := Models.GetTop(ids)
	if err != nil {
		log.Panic("dbTopList faild --- " + err.Error())
		return
	}
	utilGin.Response(http.StatusOK, "获取成功", dbDeatil)

}

func DbMusicTopList(context *gin.Context) {
	//db, err := gorm.Open("mysql", Config.DSN)
	//defer db.Close()
	//if err != nil {
	//	log.Panic("mysql db connect faild --- " + err.Error())
	//}
	//var dbMusicTopList = []Response.DbMusicTopResponse{}
	//db.Where("id >?", 0).Find(&dbMusicTopList).Scan(&dbMusicTopList)
	PageNum := 1
	PageSize := 10
	maps := make(map[string]interface{})
	dbMusicTopList, err := Models.GetMusicTopLists(PageNum, PageSize, maps)
	if err != nil {
		log.Panic("dbMusicTopList faild --- " + err.Error())
		return
	}
	utilGin := Until.Gin{Ctx: context}
	utilGin.Response(http.StatusOK, "获取成功", dbMusicTopList)
}

func DbMusicTopDetail(context *gin.Context) {
	utilGin := Until.Gin{Ctx: context}
	id, ok := context.GetQuery("id")
	if !ok {
		utilGin.Response(http.StatusBadRequest, "id is required", nil)
		return
	}
	ids, _ := strconv.Atoi(id)
	//var dbMusicDeatil = []Response.DbMusicTopResponse{}
	//db, err := gorm.Open("mysql", Config.DSN)
	//defer db.Close()
	//if err != nil {
	//	log.Panic("mysql db connect faild --- " + err.Error())
	//}
	//db.Where("id =?", id).Find(&dbMusicDeatil).Scan(&dbMusicDeatil)
	dbMusicDeatil, err := Models.GetMusicTop(ids)
	if err != nil {
		log.Panic("dbMusicDeatil faild --- " + err.Error())
		return
	}
	utilGin.Response(http.StatusOK, "获取成功", dbMusicDeatil)
}

