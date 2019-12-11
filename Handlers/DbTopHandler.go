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
	"hhxApi/Models"
	"hhxApi/Until"
	"log"
)

func GetDbTopList(context *gin.Context) {
	db, err := gorm.Open("mysql", Config.DSN)
	defer db.Close()
	if err != nil {
		log.Panic("mysql db connect faild --- " + err.Error())
	}
	var dbTopList = []Models.DbTop{}
	db.Where("id >?", 0).Find(&dbTopList).Scan(&dbTopList)
	utilGin := Until.Gin{Ctx:context}
	utilGin.Response(Config.STATUS_OK, "获取成功", dbTopList)
}
