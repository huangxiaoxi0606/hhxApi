/*
@Time : 2019/12/12 9:02
@Author : hhx06
@File : DbMusicTopModel
@Software: GoLand
*/
package Models

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

func (DbMusicTop) TableName() string {
	return "db_music_tops"
}
