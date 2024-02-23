package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lolo35/dispatch_manele/database/models"
	"github.com/lolo35/dispatch_manele/env"
	"github.com/lolo35/dispatch_manele/logger"
)

func UpdateTargetQuantity(ctx *gin.Context) {
	type validation struct {
		Lines     string `json:"lines" form:"lines" binding:"required"`
		StartDate string `json:"start_date" form:"start_date" binding:"required"`
		EndDate   string `json:"end_date" form:"end_date" binding:"required"`
	}

	var form validation

	if err := ctx.Bind(&form); err != nil {
		logger.Err(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	//var dates []string = []string{}

	startDate, err := time.Parse("2006-01-02T15:04", form.StartDate)

	if err != nil {
		logger.Err(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	endDate, err := time.Parse("2006-01-02T15:04", form.EndDate)

	if err != nil {
		logger.Err(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// for {
	// 	if startDate.Equal(endDate) {
	// 		break
	// 	}

	// 	dates = append(dates, startDate.Format("2006-01-02 15:04"))
	// 	startDate = startDate.AddDate(0, 0, 1)
	// }

	//fmt.Println(dates)

	lines := strings.Split(strings.ReplaceAll(form.Lines, "\r\n", "\n"), "\n")
	//fmt.Println(lines)
	wg := &sync.WaitGroup{}
	wg.Add(len(lines))
	for _, line := range lines {

		go func(line string) {
			defer wg.Done()
			lineid, err := FetchLineID(line)

			if err != nil {
				logger.Err(err.Error())
			}

			if err == nil {
				//wg2 := &sync.WaitGroup{}
				//wg2.Add(len(dates))
				//for _, date := range dates {

				//go func(date string, lineid int, linecode string) {
				//defer wg2.Done()
				// en, err := time.Parse("2006-01-02 15:04", date)
				// if err != nil {
				// 	logger.Err(err.Error())
				// 	return
				// }
				// end := en.AddDate(0, 0, 1)

				// if end.After(endDate) {
				// 	logger.Err("end date is after end date")
				// 	return
				// }

				// end_date := end.Format("2006-01-02 15:04")
				pitches, err := FetchPitches(lineid, startDate.Format("2006-01-02 15:04"), endDate.Format("2006-01-02 15:04"))

				if err != nil {
					logger.Err(err.Error())
					return
				}

				if len(pitches) > 0 {
					wg3 := &sync.WaitGroup{}
					wg3.Add(len(pitches))
					//fmt.Println("Date: ", date, " - ", end_date, " - ", len(pitches), " pitches", "Line: ", lineid, " - ", line, " - ", len(lines), " lines")
					for _, pitch := range pitches {
						go func(pitch models.Pitch, linecode string) {
							defer wg3.Done()
							//fmt.Println(pitch)
							part, err := FetchPart(pitch.ActualProduct)

							if err != nil {
								logger.Err(err.Error())
								return
							}

							//fmt.Println(part)

							status, err := UpdatePitch(pitch.PitchStart, pitch.PitchEnd, linecode, part.Code)

							if err != nil {
								logger.Err(err.Error())
								return
							}

							if !status.Success {
								for i := 0; i < 10; i++ {
									status, err = UpdatePitch(pitch.PitchStart, pitch.PitchEnd, linecode, part.Code)
									if err != nil {
										logger.Err(err.Error())
										fmt.Println(err.Error())
									}
									if status.Success {
										break
									}
								}
							}

							fmt.Println("Pitch start: ", pitch.PitchStart, " - Pitch end: ", pitch.PitchEnd, " - Line: ", linecode, " - Part: ", part.Code, " - Status: ", status.Success)
						}(pitch, line)
					}
					wg3.Wait()
				}
				//}(date, lineid, line)
				//}
				//wg2.Wait()
			}

		}(line)
	}

	wg.Wait()

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Target quantity updated successfully",
	})
}

func FetchLineID(line string) (int, error) {
	endpoint := fmt.Sprintf("%slines/?auth=%s&site=%s&code=%s&fields=id,code", env.Env("L2L_URL"), env.Env("L2L_AUTH"), env.Env("L2L_SITE"), line)

	method := "GET"

	client := &http.Client{}
	request, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		return 0, err
	}

	response, err := client.Do(request)
	if err != nil {
		return 0, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return 0, err

	}

	var lineResp models.LineRequest

	json.Unmarshal(body, &lineResp)

	if lineResp.Success && len(lineResp.Line) > 0 {
		return lineResp.Line[0].ID, nil
	}

	return 0, fmt.Errorf("line not found")

}

func FetchPitches(lineID int, start string, end string) ([]models.Pitch, error) {
	limit := 100
	offset := 0
	var pitchData []models.Pitch
	for {
		endpoint := fmt.Sprintf("%spitches/?auth=%s&site=%s&line=%d&shift_start_date__gte=%s&shift_start_date__lte=%s&limit=%d&offset=%d&active=true&fields=actual_product,id,pitch_start,pitch_end",
			env.Env("L2L_URL"),
			env.Env("L2L_AUTH"),
			env.Env("L2L_SITE"),
			lineID,
			url.QueryEscape(start),
			url.QueryEscape(end),
			limit,
			offset,
		)
		//fmt.Println("Start: ", start, "End: ", end, "Line: ", lineID, "Limit", limit, "Offset: ", offset, "Endpoint: ", endpoint)
		request, err := http.NewRequest("GET", endpoint, nil)

		if err != nil {
			return nil, err
		}

		client := &http.Client{}
		response, err := client.Do(request)

		if err != nil {
			return nil, err
		}

		defer response.Body.Close()

		//fmt.Println("response Status:", response.Status)

		var pitchResp models.PitchResponse
		json.NewDecoder(response.Body).Decode(&pitchResp)
		//fmt.Println("response Body:", pitchResp)

		pitchData = append(pitchData, pitchResp.Pitches...)
		fmt.Println(len(pitchResp.Pitches))
		if len(pitchResp.Pitches) < limit {
			break
		}

		offset += limit
	}

	return pitchData, nil
}

func FetchPart(partID int) (models.Part, error) {
	endpoint := fmt.Sprintf("%sproductcomponents/?auth=%s&site=%s&id=%d&fields=code", env.Env("L2L_URL"), env.Env("L2L_AUTH"), env.Env("L2L_SITE"), partID)

	request, err := http.NewRequest("GET", endpoint, nil)
	var part models.Part
	if err != nil {
		return part, err
	}

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		return part, err
	}

	defer response.Body.Close()

	var partResp models.PartResponse

	json.NewDecoder(response.Body).Decode(&partResp)

	//fmt.Println("response Body:", partResp)

	part = partResp.Parts
	return part, nil
}

func UpdatePitch(start, end, linecode, partcode string) (models.UpdatePitchResponse, error) {
	st, err := time.Parse("2006-01-02T15:04:05", start)

	if err != nil {
		return models.UpdatePitchResponse{}, err
	}

	en, err := time.Parse("2006-01-02T15:04:05", end)

	if err != nil {
		return models.UpdatePitchResponse{}, err
	}
	endpoint := fmt.Sprintf("%spitchdetails/record_details/?auth=%s&site=%s&start=%s&end=%s&linecode=%s&productcode=%s&actual=0",
		env.Env("L2L_URL"),
		env.Env("L2L_AUTH"),
		env.Env("L2L_SITE"),
		url.QueryEscape(st.Format("2006-01-02 15:04:05")),
		url.QueryEscape(en.Format("2006-01-02 15:04:05")),
		url.QueryEscape(linecode),
		url.QueryEscape(partcode),
	)
	var resp models.UpdatePitchResponse
	request, err := http.NewRequest("POST", endpoint, nil)

	if err != nil {
		return resp, err
	}

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		return resp, err
	}

	defer response.Body.Close()

	json.NewDecoder(response.Body).Decode(&resp)

	return resp, nil
}
