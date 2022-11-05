package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"parkeradam.dev/dmc-convert/controllers"
)

type rgbInput struct {
	Values []string `json:"values"`
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	controllers.InitValues(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run()
}
