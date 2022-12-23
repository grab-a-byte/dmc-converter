package main

import (
	"fmt"
	"image/color"
	"math"
	"os"
	"strings"

	"parkeradam.dev/dmc-convert/dmc"
	"parkeradam.dev/dmc-convert/rgbcolor"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readlines(path string) []string {
	data, err := os.ReadFile(path)
	check(err)
	str := string(data)
	content := strings.Split(strings.ReplaceAll(str, "\r\n", "\n"), "\n")
	return content
}

func getClosest(dmcCols []dmc.Dmc, col color.RGBA) (currentDmc dmc.Dmc) {

	var lowest = math.MaxInt32

	for _, item := range dmcCols {
		diff := rgbcolor.Compare(item.Color, col)
		if diff < lowest {
			lowest = diff
			currentDmc = item
		}
	}

	return currentDmc
}

func main() {
	dmcLines := readlines("./dmc_colors.csv")
	var dmcCols []dmc.Dmc

	for _, item := range dmcLines {
		dmcCols = append(dmcCols, dmc.FromCsvLine(item))
	}

	colorList := readlines("./colors.txt")
	colorType := strings.ToLower(colorList[0])
	colorsToConvert := colorList[1:]

	for _, item := range colorsToConvert {
		if colorType == "rgb" {
			color := rgbcolor.FromSpaceSeparatedLine(item)
			nearest := getClosest(dmcCols, color)
			fmt.Println(nearest)
		} else if colorType == "hex" {
			color, _ := rgbcolor.ParseHexColor(item)
			nearest := getClosest(dmcCols, color)
			fmt.Println(nearest)
		} else {
			panic("Unknown input type")
		}
	}
}
