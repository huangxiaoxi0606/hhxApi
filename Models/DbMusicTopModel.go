/*
@Time : 2019/12/12 9:02
@Author : hhx06
@File : DbMusicTopModel
@Software: GoLand
*/
package Models

import (
	"github.com/jinzhu/gorm"
	"hhxApi/Response"
)

type DbMusicTop struct {
	Id       int64  `json:"id"`
	No       string `json:"no"`
	Img      string `json:"img"`
	Title    string `json:"title"`
	SingName string `json:"sing_name"`
	Date     string `json:"date"`
	Album    string `json:"album"`
	Cd       string `json:"cd"`
	Type     string `json:"type"`
	Star     string `json:"star"`
	Comment  string `json:"comment"`
	Intro    string `json:"intro"`
	Songs    string `json:"songs"`
}

func GetMusicTopLists(pageNum int, pageSize int, maps interface{}) ([]*Response.DbMusicTopResponse, error) {
	var dbMusicTop []*DbMusicTop
	var dbMusicTopList = []*Response.DbMusicTopResponse{}
	errs := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&dbMusicTop).Scan(&dbMusicTopList).Error
	if errs != nil && errs != gorm.ErrRecordNotFound {
		return nil, errs
	}

	return dbMusicTopList, nil
}


func GetMusicTop(id int) (Response.DbMusicTopResponse, error) {
	var dbMusicTop DbMusicTop
	var dbMusicTopDetail = Response.DbMusicTopResponse{}
	db.Where("id = ?", id).First(&dbMusicTop).Scan(&dbMusicTopDetail)
	return dbMusicTopDetail, nil
}
