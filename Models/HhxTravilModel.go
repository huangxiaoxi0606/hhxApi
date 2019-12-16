/*
@Time : 2019/12/16 9:24
@Author : hhx06
@File : HhxTravilModel
@Software: GoLand
*/
package Models

import (
	"github.com/jinzhu/gorm"
	"hhxApi/Response"
)

type HhxTravil struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Topic        string `json:"topic"`
	Money        string `json:"money"`
	Days         int64  `json:"days"`
	Nums         int64  `json:"nums"`
	Status       string `json:"status"`
	TravilStart string `json:"travil_start"`
	TravilEnd   string `json:"travil_end"`
	Note         string `json:"note"`

}

func GetTravilLists(pageNum int, pageSize int, maps interface{}) ([]*Response.HhxTravilResponse, error) {
	var hhxTravil []*HhxTravil
	var hhxTravilList = []*Response.HhxTravilResponse{}
	errs := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&hhxTravil).Scan(&hhxTravilList).Error
	if errs != nil && errs != gorm.ErrRecordNotFound {
		return nil, errs
	}

	return hhxTravilList, nil
}

//func GetTop(id int) (Response.DbTopResponse, error) {
//	var dbtop DbTop
//	var dbTopDetail = Response.DbTopResponse{}
//	db.Where("id = ?", id).First(&dbtop).Scan(&dbTopDetail)
//	return dbTopDetail, nil
//}

func GetHhxTravilDetail(id int)(Response.HhxTravilResponse,error)  {
	var hhxTravil HhxTravil
   	var travilDetail = Response.HhxTravilResponse{}
	db.Where("id = ?", id).First(&hhxTravil).Scan(&travilDetail)
	return travilDetail, nil
}

func GetHhxTravilTraffic(id int)([]*Response.HhxTravilTrafficResponse,error)  {
	var travilTraffic = []*Response.HhxTravilTrafficResponse{}
	db.Table("travil_traffic").Where("hhx_travil_id =?",id).Select("name,money,img,illustrate").Scan(&travilTraffic)
	return travilTraffic,nil
}

func GetHhxTravilBill(id int)([]*Response.HhxTravilBillResponse,error)  {
	var travilBill = []*Response.HhxTravilBillResponse{}
	db.Table("direction_logs").Select("direction_logs.illustration,direction_logs.money").Joins("left join travil_bills on travil_bills.direction_id = direction_logs.id").Where("travil_bills.hhx_travil_id = ?",id).Scan(&travilBill)
	return travilBill, nil

}