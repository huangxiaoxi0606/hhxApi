/*
@Time : 2019/12/13 14:35
@Author : hhx06
@File : DirectionModel
@Software: GoLand
*/
package Models

import (
	"github.com/jinzhu/gorm"
	"hhxApi/Config"
	"hhxApi/Response"
	"log"
)

type Direction struct {
	Id           int64  `json:"id"`
	Name  string `json:"name"`
	Intro  string `json:"intro"`
	Img  string `json:"Img"`
	Status  string `json:"status"`
	OrderNum  int64 `json:"order_num"`
	AllNum  string `json:"all_num"`
	CreatedAt  string `json:"created_at"`
}

func GetDirectionLists(pageNum int, pageSize int, maps interface{}) ([]*Response.DirectionResponse, error) {
	var direction []*Direction
	var directionList = []*Response.DirectionResponse{}
	db, err := gorm.Open("mysql", Config.DSN)
	defer db.Close()
	if err != nil {
		log.Panic("mysql db connect faild --- " + err.Error())
	}
	errs := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&direction).Scan(&directionList).Error
	if errs != nil && errs != gorm.ErrRecordNotFound {
		return nil, errs
	}

	return directionList, nil
}