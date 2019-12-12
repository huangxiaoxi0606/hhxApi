/*
@Time : 2019/12/11 14:30
@Author : hhx06
@File : DbTopHandler
@Software: GoLand
*/
package Handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"hhxApi/Config"
	"hhxApi/Response"
	"hhxApi/Until"
	"log"
)

func GetDbTopList(context *gin.Context) {
	db, err := gorm.Open("mysql", Config.DSN)
	defer db.Close()
	if err != nil {
		log.Panic("mysql db connect faild --- " + err.Error())
	}
	var dbTopList = []Response.DbTopResponse{}
	db.Where("id >?", 0).Find(&dbTopList).Scan(&dbTopList)
	utilGin := Until.Gin{Ctx: context}
	utilGin.Response(Config.STATUS_OK, "获取成功", dbTopList)
}

func GetDbTopDetail(context *gin.Context) {
	utilGin := Until.Gin{Ctx: context}
	id, ok := context.GetQuery("id")
	if !ok {
		utilGin.Response(Config.STATUS_LACK_PARAM, "id is required", nil)
		return
	}
	var dbDeatil = []Response.DbTopResponse{}
	db, err := gorm.Open("mysql", Config.DSN)
	defer db.Close()
	if err != nil {
		log.Panic("mysql db connect faild --- " + err.Error())
	}
	db.Where("id =?", id).Find(&dbDeatil).Scan(&dbDeatil)
	utilGin.Response(Config.STATUS_OK, "获取成功", dbDeatil)

}

func GetDbMusicTopList(context *gin.Context) {
	db, err := gorm.Open("mysql", Config.DSN)
	defer db.Close()
	if err != nil {
		log.Panic("mysql db connect faild --- " + err.Error())
	}
	var dbMusicTopList = []Response.DbMusicTopResponse{}
	db.Where("id >?", 0).Find(&dbMusicTopList).Scan(&dbMusicTopList)
	utilGin := Until.Gin{Ctx: context}
	utilGin.Response(Config.STATUS_OK, "获取成功", dbMusicTopList)
}

func GetDbMusicTopDetail(context *gin.Context) {
	utilGin := Until.Gin{Ctx: context}
	id, ok := context.GetQuery("id")
	if !ok {
		utilGin.Response(Config.STATUS_LACK_PARAM, "id is required", nil)
		return
	}
	var dbMusicDeatil = []Response.DbMusicTopResponse{}
	db, err := gorm.Open("mysql", Config.DSN)
	defer db.Close()
	if err != nil {
		log.Panic("mysql db connect faild --- " + err.Error())
	}
	db.Where("id =?", id).Find(&dbMusicDeatil).Scan(&dbMusicDeatil)
	utilGin.Response(Config.STATUS_OK, "获取成功", dbMusicDeatil)

}
