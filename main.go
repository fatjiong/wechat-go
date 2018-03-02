package main

import (
	"github.com/fatjiong/wechat-go/wechat"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/", wechat.DoHandler)
	router.Run(":8888")
}
