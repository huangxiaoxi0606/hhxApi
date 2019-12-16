/*
@Time : 2019/12/16 12:15
@Author : hhx06
@File : HaveHandler
@Software: GoLand
*/
package Handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"hhxApi/Config"
	"io"
	"log"
	"net/http"
	"os"

	"hhxApi/Models"
	"hhxApi/Until"
	"time"
)

func AddHave(context *gin.Context)  {
	utilGin := Until.Gin{Ctx: context}
	//name, ok := context.GetQuery("name")
	//if !ok {
	//	utilGin.Response(Config.STATUS_LACK_PARAM, "name is required", nil)
	//	return
	//}
	//
	//intro, ok := context.GetQuery("intro")
	//if !ok {
	//	utilGin.Response(Config.STATUS_LACK_PARAM, "intro is required", nil)
	//	return
	//}
	//Pic, ok := context.GetQuery("Pic")
	//if !ok {
	//	utilGin.Response(Config.STATUS_LACK_PARAM, "Pic is required", nil)
	//	return
	//}
	//pic, ok := context.GetQuery("pic")
	//if !ok {
	//	utilGin.Response(Config.STATUS_LACK_PARAM, "Pic is required", nil)
	//	return
	//}
	//year, ok := context.GetQuery("year")
	//if !ok {
	//	utilGin.Response(Config.STATUS_LACK_PARAM, "Pic is required", nil)
	//	return
	//}

	have := Models.Have{Name: "tt", Intro: "这是intro",Pic: "Pic",Year: "Year",Mold:1,Type:1, CreatedAt:time.Now()}
	var db *gorm.DB
	var err error
	db, err = gorm.Open("mysql", Config.DSN)
	if err != nil {
		log.Panic("mysql db connect faild --- " + err.Error())
	}
	db.Create(&have)
	utilGin.Response(http.StatusOK, "获取成功", "hhxTravilTraffic")

}

func uploadFile(c *gin.Context) {
	// FormFile方法会读取参数“upload”后面的文件名，返回值是一个File指针，和一个FileHeader指针，和一个err错误。
	file, header, err := c.Request.FormFile("upload")
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	// header调用Filename方法，就可以得到文件名
	filename := header.Filename
	fmt.Println(file, err, filename)

	// 创建一个文件，文件名为filename，这里的返回值out也是一个File指针
	out, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer out.Close()

	// 将file的内容拷贝到out
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	c.String(http.StatusCreated, "upload successful \n")
}
