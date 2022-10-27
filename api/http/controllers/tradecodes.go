package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lolo35/dispatch_manele/env"
	"github.com/lolo35/dispatch_manele/logger"
)

type L2LTradesResp struct {
	Success bool         `json:"success"`
	Offset  int          `json:"offset"`
	Limit   int          `json:"limit"`
	Data    []Tradecodes `json:"data"`
	Error   string       `json:"error,omitempty"`
}

type Tradecodes struct {
	Id          int    `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

func Fetchtradecodes(ctx *gin.Context) {
	endpoint := fmt.Sprintf("%strades/?auth=%s&site=%s&active=true&fields=id,code,description", env.Env("L2L_URL"), env.Env("L2L_AUTH"), env.Env("L2L_SITE"))

	response, err := http.Get(endpoint)

	if err != nil {
		logger.Err(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		logger.Err(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	//fmt.Println(string(body))

	var tradecodes L2LTradesResp
	json.Unmarshal([]byte(body), &tradecodes)

	ctx.JSON(http.StatusOK, tradecodes)
}
