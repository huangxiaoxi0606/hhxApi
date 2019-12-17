/*
@Time : 2019/12/16 12:15
@Author : hhx06
@File : HaveHandler
@Software: GoLand
*/
package Handlers

import (
	"github.com/gin-gonic/gin"
	"hhxApi/Helpers"
	"hhxApi/Models"
	"hhxApi/Until"
	"net/http"
	"strconv"
	"time"
)

func AddAsset(context *gin.Context) {
	utilGin := Until.Gin{Ctx: context}
	name, ok := context.GetPostForm("name")
	if !ok {
		utilGin.Response(http.StatusBadRequest, "name is required", nil)
		return
	}
	mold, ok := context.GetPostForm("mold")
	if !ok {
		utilGin.Response(http.StatusBadRequest, "mold is required", nil)
		return
	}
	molds, _ := strconv.Atoi(mold)

	typeS, ok := context.GetPostForm("type")
	if !ok {
		utilGin.Response(http.StatusBadRequest, "typeS is required", nil)
		return
	}
	types, _ := strconv.Atoi(typeS)

	header, err := context.FormFile("pic")
	if err != nil {
		context.String(http.StatusBadRequest, "Bad request")
		return
	}
	pic := "public/" + header.Filename
	if err := Helpers.SaveUploadedFile(header, pic); err != nil {
		// ignore
	}
	asset := Models.Asset{Name: name, Mold: molds, Pic: pic, Type: types, CreatedAt: time.Now()}
	if err := Models.PostCreateAsset(asset); err != nil {
		utilGin.Response(http.StatusCreated, "创建失败", nil)
	}
	utilGin.Response(http.StatusOK, "创建成功", asset)

}

func AssetMolds(context *gin.Context) {
	utilGin := Until.Gin{Ctx: context}
	result := Models.GetMold()
	id, ok := context.GetPostForm("id")
	if !ok {
		utilGin.Response(http.StatusOK, "获取成功", result)
	}
	ids, _ := strconv.Atoi(id)
	utilGin.Response(http.StatusOK, "获取成功", result[ids])
}

func AssetType(context *gin.Context) {
	utilGin := Until.Gin{Ctx: context}

	key, ok := context.GetQuery("key")
	if !ok {
		utilGin.Response(http.StatusBadRequest, "key is required", nil)
		return
	}
	result := Models.GetType(key)
	utilGin.Response(http.StatusOK, "获取成功", result)
}
