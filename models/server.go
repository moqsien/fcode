package models

import (
	"fcode/cnf"
	"fcode/models/fitten"
	"fcode/models/openai"
	"strings"

	"github.com/gin-gonic/gin"
)

func Serve() {
	r := gin.Default()

	// r.Use(gin.Logger())

	r.POST("/*path", func(c *gin.Context) {
		c.Set(cnf.ModelCtxKey, cnf.DefaultModel)

		if strings.HasPrefix(cnf.DefaultModel.Name, "fitten_code") {
			fitten.HandleAll(c)
		} else {
			openai.HandleAll(c)
		}
	})

	r.Run(cnf.DefaultConf.GetPort())
}
