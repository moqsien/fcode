package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

func handleChat(c *gin.Context, body *bytes.Buffer) {
	req := &LspAIReq{}
	err := json.NewDecoder(body).Decode(req)
	if err != nil {
		return
	}
	fmt.Println(c)
}
