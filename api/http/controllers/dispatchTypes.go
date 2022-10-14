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

type L2LDispatchTypesResp struct {
	Success bool           `json:"success"`
	Offset  int            `json:"offset"`
	Limit   int            `json:"limit"`
	Error   string         `json:"error,omitempty"`
	Data    []DispatchType `json:"data"`
}

type DispatchType struct {
	Id          int    `json:"id"`
	Site        int    `json:"site"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

func Fetchdispatchtypes(ctx *gin.Context) {
	endpoint := fmt.Sprintf("%sdispatchtypes/?auth=%s&site=%s&active=true&fields=id,site,code,description&limit=200", env.Env("L2L_URL"), env.Env("L2L_AUTH"), env.Env("L2L_SITE"))

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

	var resp L2LDispatchTypesResp
	json.Unmarshal([]byte(body), &resp)

	ctx.JSON(http.StatusOK, resp)
}
