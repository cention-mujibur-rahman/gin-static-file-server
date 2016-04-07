package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("css", false)))
	r.Use(static.Serve("/js", static.LocalFile("js", false)))
	r.Use(static.Serve("/fonts", static.LocalFile("fonts", false)))
	// Listen and Server in 0.0.0.0:8080
	r.LoadHTMLGlob("html/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})
	r.Run(":8080")
}
