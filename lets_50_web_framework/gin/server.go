package main

import (
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// 获取一个默认路由 router
	r := gin.Default()
	// curl http://localhost:8080/
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"os":         runtime.GOOS,
			"go-version": runtime.Version(),
			"now_time":   time.Now().Format("2006-01-02 15:04:05"),
		})
	})
	// curl http://localhost:8080/user/ding/222
	r.GET("/user/:name/*action", func(c *gin.Context) {
		c.String(200, c.Param("name")+" / "+c.ClientIP())
	})
	// curl -i http://localhost:8080/user/ding\?id\=999
	r.GET("/user/:name", func(c *gin.Context) {
		c.String(200, c.Param("name")+" / ID="+c.DefaultQuery("id", "1"))
	})
	//fmt.Println(runtime.GOOS)
	r.Run(":8080")
}
