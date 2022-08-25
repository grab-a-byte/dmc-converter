package dmc

import (
	"image/color"
	"strconv"
	"strings"

	"parkeradam.dev/dmc-convert/rgbcolor"
)

type Dmc struct {
	Floss int
	Name  string
	Color color.RGBA
	Hex   string
}

func trimAndConvertToUInt8(str string) uint8 {
	val, _ := strconv.ParseInt(strings.TrimSpace(str), 10, 0)
	return uint8(val)
}
func trimAndConvertToInt(str string) int {
	val, _ := strconv.ParseInt(strings.TrimSpace(str), 10, 0)
	return int(val)
}

func FromCsvLine(csvLine string) Dmc {
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
	dmc.Hex = rgbcolor.HexFromRgbColor(dmc.Color)
	return dmc
}
