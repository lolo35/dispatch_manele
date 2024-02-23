package models

import "gorm.io/gorm"

type Dispatch struct {
	ID               uint `gorm:"primaryKey"`
	Dispatchnumber   int
	Dispatchid       int
	Site             int
	Dispatchtype     int
	Dispatchtypecode string
	Machinecode      string
	Machine          int
	Linecode         string
	Line             int
	gorm.Model
}

type DeletedDispatches struct {
	ID                              uint `gorm:"primaryKey"`
	Dispatchid                      int
	Reported                        string
	Currentstatus_dispatchstatus_id int
	Currentstatus_description       string
	Wonumber                        string
	Site                            int
	Dispatchtype                    int
	Dispatchtypecode                string
	Machinecode                     string
	Machine                         int
	Linecode                        string
	Line                            int
	gorm.Model
}

type DispatchDescriptions struct {
	ID                              uint `gorm:"primaryKey"`
	Dispatchtypecode                string
	Dispatchid                      int `gorm:"unique"`
	Description                     string
	Currentstatus_dispatchstatus_id int
	Currentstatus_description       string
	Wonumber                        string
	gorm.Model
}

type L2LDispatchSystemStatusRequest struct {
	Success bool                          `json:"success"`
	Data    []L2LDispatchSystemStatusData `json:"data"`
	Error   string                        `json:"error,omitempty"`
}

type L2LDispatchSystemStatusData struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}
