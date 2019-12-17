/*
@Time : 2019/12/16 12:08
@Author : hhx06
@File : AssetModel
@Software: GoLand
*/
package Models

import (
	"time"
)

type Asset struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Pic       string    `json:"pic"`
	Mold      int       `json:"mold"`
	Type      int       `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

var Mold = []string{"外套", "卫衣", "衬衣", "毛衣", "大衣", "羽绒服", "短袖", "长袖", "秋衣", "功能上衣", "牛仔裤", "休闲裤", "短裤", "运动裤", "秋裤", "毛裤", "功能裤", "功能鞋", "帆布鞋", "板鞋", "雪地靴", "拖鞋", "休闲鞋", "马丁靴", "围巾", "帽子", "睡衣", "背包", "配饰", "Music相关", "电脑相关", "手机相关", "edc", "相机相关", "coffee相关"}

//var Type = map[string][]string{"clothes": {"外套","卫衣","衬衣","毛衣","大衣","羽绒服","短袖","长袖","秋衣","功能上衣", "牛仔裤","休闲裤","短裤","运动裤","秋裤","毛裤","功能裤"}, "shoes": {"功能鞋","帆布鞋","板鞋","雪地靴","拖鞋","休闲鞋","马丁靴"},"accessories":{"围巾","帽子","睡衣","背包","配饰"},"product":{"Music相关","电脑相关","手机相关","edc","相机相关"},"coffee":{"coffee相关"}}
var Type = map[string][]int{"clothes": {0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, "shoes": {17, 18, 19, 20, 21, 22, 23}, "accessories": {24, 25, 26, 32, 34}, "product": {27, 28, 29, 30, 31}, "coffee": {33}}

func PostCreateAsset(asset Asset) error {
	if err := db.Create(&asset).Error; err != nil {
		return err
	}
	return nil
}

func GetMold() []string {
	return Mold
}

func GetType(key string) map[int]string {
	data := map[int]string{}
	for _, v := range Type[key] {
		data[v] = Mold[v]
	}
	return data
}
