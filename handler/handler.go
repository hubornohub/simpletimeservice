package handler

import (
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Timestamp string `json:"timestamp"`
	IP        string `json:"ip"`
}

func getClientIP(c *gin.Context) string {
	ip, _, err := net.SplitHostPort(c.Request.RemoteAddr)
	if err != nil {
		return c.Request.RemoteAddr
	}
	return ip
}

func GetTimeHandler(c *gin.Context) {
	resp := Response{
		Timestamp: time.Now().Format(time.RFC3339),
		IP:        getClientIP(c),
	}
	c.JSON(http.StatusOK, resp)
}
