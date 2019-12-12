/*
@Time : 2019/12/12 9:28
@Author : hhx06
@File : DirectionLogModel
@Software: GoLand
*/
package Models

type DirectionLog struct {
	Id           int64 `json:"id"`
	DirectionId  string `json:"direction_id"`
	DailyId      int64 `json:"daily_id"`
	Status       int64 `json:"status"`
	Ok           int64 `json:"ok"`
	Illustration string `json:"illustration"`
	Money        string `json:"money"`
	WeekDay      int64 `json:"week_day"`
	CreatedAt    string `json:"created_at"`
}

func (DirectionLog) TableName() string {
	return "direction_logs"
}

