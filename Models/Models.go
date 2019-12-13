/*
@Time : 2019/12/13 17:32
@Author : hhx06
@File : Models
@Software: GoLand
*/
package Models

import (
	"github.com/jinzhu/gorm"
	"hhxApi/Config"
	"log"
)

var err error
var db *gorm.DB
func Su(){
	db, err = gorm.Open("mysql", Config.DSN)
	//defer db.Close()
	if err != nil {
		log.Panic("mysql db connect faild --- " + err.Error())
	}
}