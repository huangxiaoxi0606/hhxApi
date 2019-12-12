/*
@Time : 2019/12/11 14:50
@Author : hhx06
@File : DbTop
@Software: GoLand
*/
package Models

import (
	"github.com/jinzhu/gorm"
	"hhxApi/Config"
	"hhxApi/Response"
	"log"
)

type DbTop struct {
	Id           int64  `json:"id"`
	No           string `json:"no"`
	Img          string `json:"img"`
	Ctitle       string `json:"c_title"`
	Wtitle       string `json:"w_title"`
	RatingNum    string `json:"rating_num"`
	Inq          string `json:"inq"`
	CommentNum   string `json:"comment_num"`
	Url          string `json:"url"`
	Director     string `json:"director"`
	ScreenWriter string `json:"screen_writer"`
	Actor        string `json:"actor"`
	Type         string `json:"type"`
	TimeLong     string `json:"time_long"`
	ReleaseDate  string `json:"release_date"`
	Intro        string `json:"intro"`
	Year         string `json:"year"`
	PanUrl       string `json:"pan_url"`
	PanCode      string `json:"pan_code"`
}


func ExistDbTopByID(id int) (bool, error) {
	db, err := gorm.Open("mysql", Config.DSN)
	defer db.Close()
	if err != nil {
		log.Panic("mysql db connect faild --- " + err.Error())
	}

	var dbTop DbTop
	errs := db.Select("id").Where("id = ?", id).First(&dbTop).Error
	if errs != nil && errs != gorm.ErrRecordNotFound {
		return false, errs
	}

	if dbTop.Id > 0 {
		return true, nil
	}
	return false, nil
}

func GetTopLists(pageNum int, pageSize int, maps interface{}) ([]*Response.DbTopResponse, error) {
	var dbtops []*DbTop
	var dbTopList = []*Response.DbTopResponse{}
	db, err := gorm.Open("mysql", Config.DSN)
	defer db.Close()
	if err != nil {
		log.Panic("mysql db connect faild --- " + err.Error())
	}
	errs := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&dbtops).Scan(&dbTopList).Error
	if errs != nil && errs != gorm.ErrRecordNotFound {
		return nil, errs
	}

	return dbTopList, nil
}


func GetTop(id int) (Response.DbTopResponse, error) {
	var dbtop DbTop
	var dbTopDetail = Response.DbTopResponse{}
	db, err := gorm.Open("mysql", Config.DSN)
	defer db.Close()
	if err != nil {
		log.Panic("mysql db connect faild --- " + err.Error())
	}
	db.Where("id = ?", id).First(&dbtop).Scan(&dbTopDetail)
	return dbTopDetail, nil
}