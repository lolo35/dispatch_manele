package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lolo35/dispatch_manele/database/models"
	"github.com/lolo35/dispatch_manele/env"
	"github.com/lolo35/dispatch_manele/logger"
)

func FetchDispatchStatuses(ctx *gin.Context) {
	endpoint := fmt.Sprintf("%sdispatchstatuses/get_system_statuses/?auth=%s&site=%s",
		env.Env("L2L_URL"),
		env.Env("L2L_AUTH"),
		env.Env("L2L_SITE"),
	)

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
	fmt.Println(string(body))
	var systemstatuses models.L2LDispatchSystemStatusRequest

	json.Unmarshal([]byte(body), &systemstatuses)

	if systemstatuses.Success {
		ctx.JSON(http.StatusOK, gin.H{
			"success":  true,
			"statuses": systemstatuses.Data,
		})
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"success": false,
		"error":   systemstatuses.Error,
	})
}
