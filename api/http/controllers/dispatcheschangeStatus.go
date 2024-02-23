package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/lolo35/dispatch_manele/env"
	"github.com/lolo35/dispatch_manele/logger"
)

func ChangeDispatchStatus(ctx *gin.Context) {
	type validation struct {
		Dispatchnumbers string `json:"dispatchnumbers" form:"dispatchnumbers" binding:"required"`
		Status          int    `json:"status" form:"status" binding:"required"`
	}

	var validate validation

	if err := ctx.ShouldBind(&validate); err != nil {
		logger.Err(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	dispatchnumbers := strings.Split(strings.ReplaceAll(validate.Dispatchnumbers, "\r\n", "\n"), "\n")
	var wg = &sync.WaitGroup{}
	wg.Add(len(dispatchnumbers))
	for _, dispatchnr := range dispatchnumbers {
		go func(value string, status int) {
			fmt.Println(value)
			dispatch, err := FetchDispatch(value)
			if err != nil {
				logger.Err(err.Error())
				ctx.JSON(http.StatusBadRequest, gin.H{
					"success": false,
					"error":   err.Error(),
				})
				return
			}
			//fmt.Println(dispatch.Success)
			//fmt.Println(dispatch)
			if len(dispatch.Data) > 0 {
				endpoint := fmt.Sprintf("%sdispatches/changestatus/%d/?auth=%s",
					env.Env("L2L_URL"),
					dispatch.Data[0].Id,
					env.Env("L2L_AUTH"),
				)

				method := "POST"

				payload := strings.NewReader(
					fmt.Sprintf("auth=%s&site=%s&dispatchstatus_id=%d", env.Env("L2L_AUTH"), env.Env("L2L_SITE"), validate.Status))

				client := http.Client{}

				request, err := http.NewRequest(method, endpoint, payload)
				if err != nil {
					logger.Err(err.Error())
					return
				}

				request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
				res, err := client.Do(request)

				if err != nil {
					logger.Err(err.Error())
					return
				}

				defer res.Body.Close()

				body, err := io.ReadAll(res.Body)
				if err != nil {
					logger.Err(err.Error())
					return
				}

				fmt.Println(string(body))

				type response struct {
					Success bool   `json:"success"`
					Error   string `json:"error,omitempty"`
				}

				var l2lresp response
				json.Unmarshal([]byte(body), &l2lresp)
				if !l2lresp.Success {
					logger.Err(l2lresp.Error)
					return
				}
				if status == -6 {
					CloseDispatch(dispatch.Data[0].Id)
				}
			}
			defer wg.Done()
		}(dispatchnr, validate.Status)
	}
	wg.Wait()
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func CloseDispatch(dispatch_id int) (bool, error) {
	endpoint := fmt.Sprintf("%sdispatches/close/%d/?auth=%s&site=%s", env.Env("L2L_URL"), dispatch_id, env.Env("L2L_AUTH"), env.Env("L2L_SITE"))
	response, err := http.Get(endpoint)

	if err != nil {
		return false, err
	}
	defer response.Body.Close()

	type l2lresp struct {
		Success bool `json:"success"`
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		logger.Err(err.Error())
		return false, err
	}

	var resp l2lresp
	json.Unmarshal([]byte(body), &resp)

	if resp.Success {
		return true, nil
	}

	return false, errors.New("error closing dispatch")
}
