/*
@Time : 2019/12/16 12:08
@Author : hhx06
@File : HaveModel
@Software: GoLand
*/
package Models

import "time"

type Have struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Intro     string `json:"intro"`
	Pic       string `json:"pic"`
	Year      string `json:"year"`
	Mold      int64  `json:"mold"`
	Type      int64  `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}


