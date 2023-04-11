package controllers

import (
	"image/color"
	"net/http"

	"github.com/gin-gonic/gin"
	"parkeradam.dev/dmc-convert/dmc"
	"parkeradam.dev/dmc-convert/extendedcolor"
)

func InitValues(engine *gin.Engine) {
	engine.GET("/dmc/get", func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, dmc.Get())
	})
	engine.POST("/dmc/submit", translateValues)
}

type ValueInput struct {
	Inputs    []string `json:"inputs"`
	InputType string   `json:"inputType"`
}

func translateValues(context *gin.Context) {
	var body ValueInput
	if err := context.BindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, "invalid body")
	}
	var rgbColors []color.RGBA
	if body.InputType == "hex" {
		rgbColors = extendedcolor.MapManyHex(body.Inputs)
	} else if body.InputType == "csv" {
		rgbColors = extendedcolor.MapCommaSeperatedLinesToColors(body.Inputs)
	} else if body.InputType == "space" {
		rgbColors = extendedcolor.MapSpaceSeperatedLinesToColors(body.Inputs)
	} else {
		context.IndentedJSON(http.StatusBadRequest, "Request input type unknown")
		return
	}

	values := parseValues(rgbColors)
	context.IndentedJSON(http.StatusAccepted, values)
}

func parseValues(values []color.RGBA) []dmc.Dmc {
	var arr []dmc.Dmc
	for _, val := range values {
		closest := dmc.GetClosest(val)
		arr = append(arr, closest)
	}
	return arr
}
