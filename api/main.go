package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lolo35/dispatch_manele/database"
	"github.com/lolo35/dispatch_manele/env"
	"github.com/lolo35/dispatch_manele/http/controllers"
	"github.com/lolo35/dispatch_manele/logger"
)

func main() {
	if dberr := database.CreateDBConnection(); dberr != nil {
		logger.Err(dberr.Error())
		log.Fatal(dberr.Error())
	}
	err := database.Migrate()
	if err != nil {
		logger.Err(err.Error())
		log.Fatal(err.Error())
	}

	port := env.Env("PORT")

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	v1 := r.Group("api/v1")
	{
		v1.GET("/dispatchtypes", controllers.Fetchdispatchtypes)
		v1.GET("/lines", controllers.FetchLines)
		v1.GET("/tradecodes", controllers.Fetchtradecodes)
		v1.POST("/addDispatch", controllers.AddDispatches)
		v1.POST("/delete", controllers.DeleteDispatch)
		v1.POST("/save_description", controllers.SaveDispatchDescriptions)
		v1.GET("/description_count", controllers.FetchDispatchDescriptionCount)
		v1.GET("/dispatch_system_statuses", controllers.FetchDispatchStatuses)
		v1.POST("/change_dispatch_status", controllers.ChangeDispatchStatus)
		v1.POST("/change_dispatch_type", controllers.ChangeDispatchTypes)
		v1.POST("/update_target_quantity", controllers.UpdateTargetQuantity)

	}

	r.Run(port)
}
