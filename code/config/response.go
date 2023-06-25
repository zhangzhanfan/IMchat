package config

/*
	封装响应信息
*/

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
}

// 封装成功的响应
func Success(data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": SUCCESS_CODE,
		"data": data,
	})
}

// 封装失败的响应
func Failed(data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": FAILED_CODE,
		"data": data,
	})
}
