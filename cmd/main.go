package main

import (
	"SimpleTimeService/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", handler.GetTimeHandler)
	r.Run(":8080")
}
