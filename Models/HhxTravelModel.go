/*
@Time : 2019/12/16 9:24
@Author : hhx06
@File : HhxTravelModel
@Software: GoLand
*/
package Models

import (
	"github.com/jinzhu/gorm"
	"hhxApi/Response"
)

type HhxTravel struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Topic        string `json:"topic"`
	Money        string `json:"money"`
	Days         int64  `json:"days"`
	Nums         int64  `json:"nums"`
	Status       string `json:"status"`
	TravelStart string `json:"travel_start"`
	TravelEnd   string `json:"travel_end"`
	Note         string `json:"note"`

}

func GetTravelLists(pageNum int, pageSize int, maps interface{}) ([]*Response.HhxTravelResponse, error) {
	var hhx []*HhxTravel
	var hhxList []*Response.HhxTravelResponse
	errs := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&hhx).Scan(&hhxList).Error
	if errs != nil && errs != gorm.ErrRecordNotFound {
		return nil, errs
	}

	return hhxList, nil
}

//func GetTop(id int) (Response.DbTopResponse, error) {
//	var dbtop DbTop
//	var dbTopDetail = Response.DbTopResponse{}
//	db.Where("id = ?", id).First(&dbtop).Scan(&dbTopDetail)
//	return dbTopDetail, nil
//}

func GetHhxTravelDetail(id int)(Response.HhxTravelResponse,error)  {
	var hhx HhxTravel
   	var intro = Response.HhxTravelResponse{}
	db.Where("id = ?", id).First(&hhx).Scan(&intro)
	return intro, nil
}

func GetHhxTravelTraffic(id int)([]*Response.HhxTravelTrafficResponse,error)  {
	var traffic []*Response.HhxTravelTrafficResponse
	db.Table("travel_traffic").Where("hhx_travel_id =?",id).Select("name,money,img,illustrate").Scan(&traffic)
	return traffic,nil
}

func GetHhxTravelBill(id int)([]*Response.HhxTravelBillResponse,error)  {
	var bill []*Response.HhxTravelBillResponse
	db.Table("direction_logs").Select("direction_logs.illustration,direction_logs.money").Joins("left join travel_bills on travel_bills.direction_id = direction_logs.id").Where("travel_bills.hhx_travel_id = ?",id).Scan(&bill)
	return bill, nil

}