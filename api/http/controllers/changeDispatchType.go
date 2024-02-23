package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/lolo35/dispatch_manele/env"
	"github.com/lolo35/dispatch_manele/logger"
)

func ChangeDispatchTypes(ctx *gin.Context) {
	type validation struct {
		Dispatchnumbers string `json:"dispatchnumbers" form:"dispatchnumbers" binding:"required"`
		Dispatchtype    int    `json:"dispatchtype" form:"dispatchtype" binding:"required"`
		IsClosed        *int   `json:"is_closed" form:"is_closed" binding:"required"`
	}

	var validate validation

	if err := ctx.Bind(&validate); err != nil {
		logger.Err(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	dispatchNumbers := strings.Split(strings.ReplaceAll(validate.Dispatchnumbers, "\r\n", "\n"), "\n")
	var wg = &sync.WaitGroup{}
	wg.Add(len(dispatchNumbers))
	for _, dispatchnr := range dispatchNumbers {
		go func(dispatchnr string) {
			//fmt.Println(dispatchnr)
			dispatch, err := FetchDispatch(dispatchnr)
			if err != nil {
				logger.Err(err.Error())
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"error":   err.Error(),
				})
				return
			}

			fmt.Println(dispatch.Success)
			if len(dispatch.Data) > 0 {
				endpoint := fmt.Sprintf("%sdispatches/changetype/%d/?auth=%s",
					env.Env("L2L_URL"),
					dispatch.Data[0].Id,
					env.Env("L2L_AUTH"),
				)
				method := "POST"
				payload := strings.NewReader(fmt.Sprintf("auth=%s&site=%s&dispatchtype_id=%d", env.Env("L2L_AUTH"), env.Env("L2L_SITE"), validate.Dispatchtype))

				client := &http.Client{}
				request, err := http.NewRequest(method, endpoint, payload)
				if err != nil {
					logger.Err(err.Error())
					// ctx.JSON(http.StatusInternalServerError, gin.H{
					// 	"success": false,
					// 	"error":   err.Error(),
					// })
					return
				}
				request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

				response, err := client.Do(request)
				if err != nil {
					logger.Err(err.Error())
					// ctx.JSON(http.StatusInternalServerError, gin.H{
					// 	"success": false,
					// 	"error":   err.Error(),
					// })
					return
				}
				defer response.Body.Close()

				body, err := io.ReadAll(response.Body)
				if err != nil {
					logger.Err(err.Error())
					// ctx.JSON(http.StatusInternalServerError, gin.H{
					// 	"success": false,
					// 	"error":   err.Error(),
					// })
					return
				}

				type resp struct {
					Success bool   `json:"success"`
					Error   string `json:"error,omitempty"`
				}

				var res resp
				json.Unmarshal(body, &res)
				if !res.Success {
					logger.Err(res.Error)
					// ctx.JSON(http.StatusInternalServerError, gin.H{
					// 	"success": false,
					// 	"error":   res.Error,
					// })
					return
				}
				if *validate.IsClosed == 1 {
					CloseDispatch(dispatch.Data[0].Id)
				}
			}
			defer wg.Done()
		}(dispatchnr)
	}
	wg.Wait()

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
