package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lolo35/dispatch_manele/env"
	"github.com/lolo35/dispatch_manele/http/controllers"
)

func main() {
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
	}

	r.Run(port)
}
