/*
@Time : 2019/12/12 11:31
@Author : hhx06
@File : Response
@Software: GoLand
*/
package Response

type DbTopResponse struct {
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

type DbMusicTopResponse struct {
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

type DirectionLogResponse struct {
	Id           int64 `json:"id"`
	Name  string `json:"name"`
	DailyId      int64 `json:"daily_id"`
	Status       int64 `json:"status"`
	Ok           int64 `json:"ok"`
	Illustration string `json:"illustration"`
	Money        string `json:"money"`
	WeekDay      int64 `json:"week_day"`
	CreatedAt    string `json:"created_at"`
}

type DirectionResponse struct {
	Id           int64  `json:"id"`
	Name  string `json:"name"`
	Intro  string `json:"intro"`
	Img  string `json:"Img"`
	Status  string `json:"status"`
	OrderNum  int64 `json:"order_num"`
	AllNum  string `json:"all_num"`
	CreatedAt  string `json:"created_at"`
}


type DirectionSumResponse struct{
	Name     string  `json:"name"`
	Total        string `json:"total"`
}