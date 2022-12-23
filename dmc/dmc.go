package dmc

import (
	_ "embed"
	"image/color"
	"math"
	"strconv"
	"strings"

	"parkeradam.dev/dmc-convert/extendedcolor"
)

//go:embed dmc_colors.csv
var dmc_colors_str string

var dmcCols []Dmc

func init() {
	var dmcLines []string = strings.Split(strings.ReplaceAll(dmc_colors_str, "\r\n", "\n"), "\n")
	for _, item := range dmcLines {
		dmcCols = append(dmcCols, fromCsvLine(item))
	}
}

func Get() []Dmc {
	return dmcCols
}

type Dmc struct {
	Floss int        `json:"floss"`
	Name  string     `json:"name"`
	Color color.RGBA `json:"color"`
	Hex   string     `json:"hex"`
}

func GetClosest(col color.RGBA) Dmc {

	var currentDmc Dmc
	lowest := math.MaxInt32

	for _, item := range dmcCols {
		diff := extendedcolor.Compare(item.Color, col)
		if diff < lowest {
			lowest = diff
			currentDmc = item
		}
	}

	return currentDmc
}

func trimAndConvertToUInt8(str string) uint8 {
	val, _ := strconv.ParseInt(strings.TrimSpace(str), 10, 0)
	return uint8(val)
}
func trimAndConvertToInt(str string) int {
	val, _ := strconv.ParseInt(strings.TrimSpace(str), 10, 0)
	return int(val)
}

func fromCsvLine(csvLine string) Dmc {
	items := strings.Split(csvLine, ",")
	dmc := Dmc{
		Floss: trimAndConvertToInt(items[0]),
		Name:  strings.TrimSpace(items[1]),
		Color: color.RGBA{
			R: trimAndConvertToUInt8(items[2]),
			G: trimAndConvertToUInt8(items[3]),
			B: trimAndConvertToUInt8(items[4]),
		},
	}
	dmc.Hex = extendedcolor.HexFromRgbColor(dmc.Color)
	return dmc
}
