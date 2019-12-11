/*
@Time : 2019/12/11 14:50
@Author : hhx06
@File : DbTop
@Software: GoLand
*/
package Models

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
func (DbTop) TableName() string {
	return "db_tops"
}