/*
@Time : 2019/12/11 10:17
@Author : hhx06
@File : TestHandler
@Software: GoLand
*/
package Handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hhxApi/Until"
	"time"
)

func AesTest(context *gin.Context) {
	startTime := time.Now()
	utilGin := Until.Gin{Ctx:context}
	utilGin.Response(1, fmt.Sprintf("%v", time.Since(startTime)), nil)
}
