package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Serve() {
	r := gin.Default()

	// r.Use(gin.Logger())

	r.POST("/*path", func(c *gin.Context) {
		handleAll(c)
	})

	r.Run(fmt.Sprintf(":%+v", DefaultConf.GetPort()))
}
