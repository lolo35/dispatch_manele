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

type L2LLinesResp struct {
	Success bool    `json:"success"`
	Offset  int     `json:"offset"`
	Limit   int     `json:"limit"`
	Data    []Lines `json:"data"`
	Error   string  `json:"error,omitempty"`
}

type Lines struct {
	Id             int    `json:"id"`
	Code           string `json:"code"`
	Areacode       string `json:"areacode"`
	Description    string `json:"description"`
	Abbreviation   string `json:"abbreviation"`
	Defaultmachine int    `json:"defaultmachine"`
}

func FetchLines(ctx *gin.Context) {
	endpoint := fmt.Sprintf("%slines/?auth=%s&site=%s&active=true&fields=id,code,areacode,area,description,defaultmachine,abbreviation&limit=500", env.Env("L2L_URL"), env.Env("L2L_AUTH"), env.Env("L2L_SITE"))

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

	var l2lresp L2LLinesResp
	json.Unmarshal([]byte(body), &l2lresp)

	ctx.JSON(http.StatusOK, l2lresp)

}
