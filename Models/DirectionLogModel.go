/*
@Time : 2019/12/12 9:28
@Author : hhx06
@File : DirectionLogModel
@Software: GoLand
*/
package Models

import (
	"hhxApi/Helpers"
	"hhxApi/Response"
)

type DirectionLog struct {
	Id           int64  `json:"id"`
	DirectionId  string `json:"direction_id"`
	DailyId      int64  `json:"daily_id"`
	Status       int64  `json:"status"`
	Ok           int64  `json:"ok"`
	Illustration string `json:"illustration"`
	Money        string `json:"money"`
	WeekDay      int64  `json:"week_day"`
	CreatedAt    string `json:"created_at"`
}

func (DirectionLog) TableName() string {
	return "direction_logs"
}

func GetLog(mold int) ([]*Response.DirectionLogResponse, error) {
	flagDate := Helpers.GetDateQuery(mold)
	result := []*Response.DirectionLogResponse{}
	db.Table("direction_logs").Select("direction_logs.illustration,direction_logs.id,"+
		"direction_logs.daily_id,direction_logs.status,direction_logs.ok,direction_logs.money,direction_logs.week_day,"+
		"direction_logs.week_day,direction_logs.created_at, directions.name").Joins("left join directions on directions.id = direction_logs.direction_id").Where(" direction_logs.created_at > ?", flagDate.Format("2006-01-02")).Scan(&result)
	return result, nil
}

func GetDirectionSum(mold int) ([]*Response.DirectionSumResponse, error) {
	flagDate := Helpers.GetDateQuery(mold)
	result := []*Response.DirectionSumResponse{}
	db.Table("direction_logs").Select("directions.name, sum(direction_logs.money) as total").Joins("join directions on direction_logs.direction_id  = directions.id").Where("direction_logs.created_at >?", flagDate.Format("2006-01-02")).Group("directions.name").Scan(&result)
	return result, nil
}

func GetDirectionLog(mold int,id int) ([]*Response.DirectionLogResponse, error) {
	flagDate := Helpers.GetDateQuery(mold)
	result := []*Response.DirectionLogResponse{}
	db.Table("direction_logs").Select("direction_logs.illustration,direction_logs.id,"+
		"direction_logs.daily_id,direction_logs.status,direction_logs.ok,direction_logs.money,direction_logs.week_day,"+
		"direction_logs.week_day,direction_logs.created_at, directions.name").Joins("left join directions on directions.id = direction_logs.direction_id").Where(" direction_logs.created_at > ? and directions.id = ?", flagDate.Format("2006-01-02"),id).Scan(&result)
	return result, nil
}

