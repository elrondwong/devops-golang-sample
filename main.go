package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloWorldHandle(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

func main() {
	r := gin.Default()
	r.GET("/", HelloWorldHandle)
	r.Run(":80") // listen and serve on 0.0.0.0:80 (for windows "localhost:80")
}
