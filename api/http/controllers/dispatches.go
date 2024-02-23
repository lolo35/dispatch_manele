package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lolo35/dispatch_manele/database"
	"github.com/lolo35/dispatch_manele/database/models"
	"github.com/lolo35/dispatch_manele/env"
	"github.com/lolo35/dispatch_manele/logger"
	"gorm.io/gorm/clause"
)

type MachineResp struct {
	Success bool     `json:"success"`
	Data    Machines `json:"data"`
	Error   string   `json:"error,omitempty"`
}

type Machines struct {
	Id             int    `json:"id"`
	Code           string `json:"code"`
	Description    string `json:"description"`
	Costcentercode string `json:"costcentercode"`
}

type AddDispatchResp struct {
	Success bool            `json:"success"`
	Data    AddDispatchData `json:"data"`
	Error   string          `json:"error,omitempty"`
}

type AddDispatchData struct {
	Id               int    `json:"id"`
	Site             int    `json:"site"`
	Dispatchnumber   int    `json:"dispatchnumber"`
	Dispatchtype     int    `json:"dispatchtype"`
	Dispatchtypecode string `json:"dispatchtypecode"`
	Machinecode      string `json:"machinecode"`
	Machine          int    `json:"machine"`
	Linecode         string `json:"linecode"`
	Line             int    `json:"line"`
}

type DispatchInfoResp struct {
	Success bool           `json:"success"`
	Offset  int            `json:"offset"`
	Limit   int            `json:"limit"`
	Data    []DispatchInfo `json:"data"`
	Error   string         `json:"error,omitempty"`
}

type DispatchInfo struct {
	Id                              int    `json:"id"`
	Description                     string `json:"description,omitempty"`
	Reported                        string `json:"reported,omitempty"`
	Currentstatus_dispatchstatus_id int    `json:"currentstatus_dispatchstatus_id,omitempty"`
	Currentstatus_description       string `json:"currentstatus_description,omitempty"`
	Wonumber                        string `json:"wonumber,omitempty"`
	Site                            int    `json:"site,omitempty"`
	Dispatchtype                    int    `json:"dispatchtype,omitempty"`
	Dispatchtypecode                string `json:"dispatchtypecode,omitempty"`
	Machinecode                     string `json:"machinecode,omitempty"`
	Machine                         int    `json:"machine,omitempty"`
	Linecode                        string `json:"linecode,omitempty"`
	Line                            int    `json:"line,omitempty"`
}

func AddDispatches(ctx *gin.Context) {
	type request struct {
		Dispatchyypecode    string `json:"dispatchtypecode" form:"dispatchtypecode" binding:"required"`
		Description         string `json:"description" form:"description"`
		Tradecode           string `json:"tradecode" form:"tradecode" binding:"required"`
		Lines               string `json:"lines" form:"lines" binding:"required"`
		Resourse            string `json:"resourse" form:"resourse"`
		RandStart           int    `json:"randstart" form:"randstart" binding:"required"`
		RandEnd             int    `json:"randend" form:"randend" binding:"required"`
		DescriptionIsRandom string `json:"descriptionIsRandom" form:"descriptionIsRandom" binding:"required"`
	}

	type lines struct {
		Id             int    `json:"id"`
		Code           string `json:"code"`
		Areacode       string `json:"areacode"`
		Description    string `json:"description"`
		Abbreviation   string `json:"abbreviation"`
		Defaultmachine int    `json:"defaultmachine"`
	}

	// endpoint := fmt.Sprintf("%sdispatches/add/", env.Env("L2L_URL"))
	var validate request

	if err := ctx.ShouldBind(&validate); err != nil {
		logger.Err(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	descriptionIsRandom, err := strconv.ParseBool(ctx.PostForm("descriptionIsRandom"))
	if err != nil {
		logger.Err(err.Error())
	}
	randStart, err := strconv.Atoi(ctx.PostForm("randstart"))
	if err != nil {
		logger.Err(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
	}
	randEnd, err := strconv.Atoi(ctx.PostForm("randend"))
	if err != nil {
		logger.Err(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
	}

	if randStart > randEnd {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Start value cannot be greater than end value",
		})
		return
	}

	linesReq := ctx.PostForm("lines")
	dispatchTypeCode := ctx.PostForm("dispatchtypecode")
	description := ctx.PostForm("description")
	tradecode := ctx.PostForm("tradecode")

	resourse := ctx.PostForm("resourse")

	var linesArr []lines
	json.Unmarshal([]byte(linesReq), &linesArr)

	for _, value := range linesArr {
		machine, err := FetchMachine(value.Defaultmachine)
		if err != nil {
			logger.Err(err.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}

		fmt.Println(description)

		derr := AddDispatch(dispatchTypeCode, description, machine.Data.Code, tradecode, resourse, randStart, randEnd, descriptionIsRandom)

		if derr != nil {
			logger.Err(derr.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   derr.Error(),
			})
			return
		}
		//logger.Log(fmt.Sprintf("Opened dispatch: %s with dispatchnumber: %d", dispatch.Data.Dispatchtypecode, dispatch.Data.Dispatchnumber))
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})

}

func FetchMachine(id int) (MachineResp, error) {
	endpoint := fmt.Sprintf("%smachines/?auth=%s&site=%s&fields=id,code,description&id=%d", env.Env("L2L_URL"), env.Env("L2L_AUTH"), env.Env("L2L_SITE"), id)

	response, err := http.Get(endpoint)
	var resp MachineResp

	if err != nil {
		logger.Err(err.Error())
		return resp, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		logger.Err(err.Error())
		return resp, err
	}

	json.Unmarshal([]byte(body), &resp)

	if !resp.Success {
		logger.Err(resp.Error)
		return resp, errors.New(resp.Error)
	}

	return resp, nil
}

func fetchRandomDescription(dispatchTypeCode string) (string, error) {
	rand.Seed(time.Now().UnixNano())
	var descriptionStart models.DispatchDescriptions
	var descriptionEnd models.DispatchDescriptions
	var dbDescription models.DispatchDescriptions

	db, connErr := database.GetDatabaseConnection()
	if connErr != nil {
		logger.Err(connErr.Error())
		return "", connErr
	}

	db.Model(&models.DispatchDescriptions{}).Where("dispatchtypecode = ?", dispatchTypeCode).First(&descriptionStart)
	db.Model(&models.DispatchDescriptions{}).Where("dispatchtypecode = ?", dispatchTypeCode).Last(&descriptionEnd)
	rand.Seed(time.Now().UnixNano())

	randomNumber := rand.Intn(int(descriptionEnd.ID)-int(descriptionStart.ID)+1) + int(descriptionStart.ID)

	db.Debug().Model(&models.DispatchDescriptions{}).Where("id = ?", randomNumber).First(&dbDescription)

	description := dbDescription.Description

	return description, nil
}

func AddDispatch(dispatchtypecode string, description string, machinecode string, tradecode string, resourse string, randStart int, randEnd int, descriptionIsRandom bool) error {
	rand.Seed(time.Now().UnixNano())
	randNr := rand.Intn(randEnd-randStart+1) + randStart
	for i := 0; i < randNr; i++ {
		endpoint := fmt.Sprintf("%sdispatches/add/?auth=%s", env.Env("L2L_URL"), env.Env("L2L_AUTH"))
		method := "POST"

		var response AddDispatchResp
		var desc string
		desc = description

		if descriptionIsRandom {
			d, err := fetchRandomDescription(dispatchtypecode)
			if err != nil {
				logger.Err(err.Error())
				return err
			}
			desc = d
		}

		if machinecode == "" {
			continue
		}

		payload := strings.NewReader(
			fmt.Sprintf("auth%s&site=%s&dispatchtypecode=%s&description=%s&machinecode=%s&tradecode=%s&resource=%s", env.Env("L2L_AUTH"), env.Env("L2L_SITE"), dispatchtypecode, desc, machinecode, tradecode, resourse))

		client := http.Client{}
		request, err := http.NewRequest(method, endpoint, payload)

		if err != nil {
			logger.Err(err.Error())
			return err
		}

		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		res, err := client.Do(request)

		if err != nil {
			logger.Err(err.Error())
			return err
		}

		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)

		if err != nil {
			logger.Err(err.Error())
			return err
		}

		//fmt.Println(string(body))
		json.Unmarshal([]byte(body), &response)

		if !response.Success {
			logger.Err(response.Error)
			return errors.New(response.Error)
		}

		db, connErr := database.GetDatabaseConnection()
		if connErr != nil {
			logger.Err(connErr.Error())
		}

		result := db.Create(&models.Dispatch{
			Dispatchnumber:   response.Data.Dispatchnumber,
			Dispatchid:       response.Data.Id,
			Site:             response.Data.Site,
			Dispatchtype:     response.Data.Dispatchtype,
			Dispatchtypecode: response.Data.Dispatchtypecode,
			Machinecode:      response.Data.Machinecode,
			Machine:          response.Data.Machine,
			Linecode:         response.Data.Linecode,
			Line:             response.Data.Line,
		})

		fmt.Println("Dispatch number :", response.Data.Dispatchnumber)
		fmt.Println(result.RowsAffected)
		if result.Error != nil && result.RowsAffected != 1 {
			logger.Err(result.Error.Error())
		}
	}

	return nil
}

func DeleteDispatch(ctx *gin.Context) {
	test := ctx.PostForm("test")

	dispatchNumbers := strings.Split(strings.ReplaceAll(test, "\r\n", "\n"), "\n")
	var wg = &sync.WaitGroup{}
	wg.Add(len(dispatchNumbers))
	for _, value := range dispatchNumbers {
		go func(val string) {
			fmt.Println("starting the process")
			if val != "" {

				dispatch, err := FetchDispatch(val)
				if err != nil {
					logger.Err(err.Error())
					ctx.JSON(http.StatusBadRequest, gin.H{
						"success": false,
						"error":   err.Error(),
					})
					return
				}
				if dispatch.Success {
					for _, value := range dispatch.Data {
						delete := DispatchDelete(value.Id)

						if delete {
							db, connErr := database.GetDatabaseConnection()
							if connErr != nil {
								logger.Err(connErr.Error())
							}
							db.Create(&models.DeletedDispatches{
								Dispatchid:                      value.Id,
								Reported:                        value.Reported,
								Currentstatus_dispatchstatus_id: value.Currentstatus_dispatchstatus_id,
								Currentstatus_description:       value.Currentstatus_description,
								Wonumber:                        value.Wonumber,
								Site:                            value.Site,
								Dispatchtype:                    value.Dispatchtype,
								Dispatchtypecode:                value.Dispatchtypecode,
								Machinecode:                     value.Machinecode,
								Machine:                         value.Machine,
								Linecode:                        value.Linecode,
								Line:                            value.Line,
							})
						}
					}
				}
			}
			defer wg.Done()
		}(value)
	}
	wg.Wait()

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func DispatchDelete(dispatchid int) bool {
	endpoint := fmt.Sprintf("%sdispatches/delete/%d?auth=%s&site=%s", env.Env("L2L_URL"), dispatchid, env.Env("L2L_AUTH"), env.Env("L2L_SITE"))

	request, err := http.Get(endpoint)

	type resp struct {
		Success bool `json:"success"`
	}

	var response resp

	if err != nil {
		logger.Err(err.Error())
		return false
	}

	defer request.Body.Close()

	body, err := io.ReadAll(request.Body)

	if err != nil {
		logger.Err(err.Error())
		return false
	}

	json.Unmarshal([]byte(body), &response)

	return response.Success
}

func FetchDispatch(dispatchnumber string) (DispatchInfoResp, error) {
	endpoint := fmt.Sprintf("%sdispatches/?auth=%s&site=%s&dispatchnumber=%s&fields=id,reported,site,dispatchtype,dispatchtypecode,machinecode,machine,linecode,line", env.Env("L2L_URL"), env.Env("L2L_AUTH"), env.Env("L2L_SITE"), dispatchnumber)
	var response DispatchInfoResp
	request, err := http.Get(endpoint)

	request.Close = true

	if err != nil {
		logger.Err(err.Error())
		return response, err
	}

	defer request.Body.Close()

	body, err := io.ReadAll(request.Body)

	if err != nil {
		logger.Log(err.Error())
		return response, err
	}

	json.Unmarshal([]byte(body), &response)

	return response, nil
}

func SaveDispatchDescriptions(ctx *gin.Context) {
	type request struct {
		Dispatchtypecode string `json:"dispatchtypecode" form:"dispatchtypecode" binding:"required"`
	}

	var validate request

	if err := ctx.ShouldBind(&validate); err != nil {
		logger.Err(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	dispatchtypecode := ctx.PostForm("dispatchtypecode")

	limit := 100
	offset := 0
	for ok := true; ok; ok = true {
		l2l_endpoint := fmt.Sprintf("%sdispatches/?auth=%s&site=%s&fields=id,description&dispatchtypecode=%s&offset=%d&limit=%d", env.Env("L2L_URL"), env.Env("L2L_AUTH"), env.Env("L2L_SITE"), dispatchtypecode, offset, limit)

		req, err := http.Get(l2l_endpoint)

		if err != nil {
			logger.Err(err.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}

		defer req.Body.Close()

		body, err := io.ReadAll(req.Body)

		if err != nil {
			logger.Err(err.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   err.Error(),
			})

			return
		}

		var info DispatchInfoResp
		json.Unmarshal([]byte(body), &info)
		var descriptions []models.DispatchDescriptions
		if info.Success {
			for _, value := range info.Data {
				dispatch := models.DispatchDescriptions{
					Dispatchtypecode:                dispatchtypecode,
					Dispatchid:                      value.Id,
					Description:                     value.Description,
					Currentstatus_dispatchstatus_id: value.Currentstatus_dispatchstatus_id,
					Currentstatus_description:       value.Currentstatus_description,
					Wonumber:                        value.Wonumber,
				}

				descriptions = append(descriptions, dispatch)
			}
			db, connErr := database.GetDatabaseConnection()
			if connErr != nil {
				logger.Err(connErr.Error())
				ctx.JSON(http.StatusBadRequest, gin.H{
					"success": false,
					"error":   connErr.Error(),
				})
				return
			}
			db.Clauses(clause.OnConflict{DoNothing: true}).Save(&descriptions)
			if len(info.Data) < limit {
				break
			}
		}
		offset += 100
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func FetchDispatchDescriptionCount(ctx *gin.Context) {
	type request struct {
		Dispatchtypecode string `json:"dispatchtypecode" form:"dispatchtypecode" binding:"required"`
	}

	var validate request
	if err := ctx.ShouldBind(&validate); err != nil {
		logger.Err(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	dispatchtypecode := ctx.Query("dispatchtypecode")
	db, connErr := database.GetDatabaseConnection()

	if connErr != nil {
		logger.Err(connErr.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   connErr.Error(),
		})
		return
	}
	var count int64
	db.Model(&models.DispatchDescriptions{}).Where("dispatchtypecode = ?", dispatchtypecode).Count(&count)

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"count":   count,
	})
}
