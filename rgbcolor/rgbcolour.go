package rgbcolor

import (
	"image/color"
	"math"
	"strconv"
	"strings"
)

func intAbs(num int) int { return int(math.Abs(float64(num))) }

func subtractHighestFromLowest(i1, i2 uint8) int {
	if i1 > i2 {
		return int(i1 - i2)
	}
	return int(i2 - i1)
}

func Compare(col1 color.RGBA, col2 color.RGBA) int {
	var diff int = subtractHighestFromLowest(col1.R, col2.R) +
		subtractHighestFromLowest(col1.G, col2.G) +
		subtractHighestFromLowest(col1.B, col2.B)

	return intAbs(diff)
}

func parseToUint8(str string) uint8 {
	value, _ := strconv.ParseUint(str, 10, 8)
	return uint8(value)
}

func FromSpaceSepeartedLine(str string) (col color.RGBA) {
	segments := strings.Split(str, " ")
	col.R = parseToUint8(segments[0])
	col.G = parseToUint8(segments[1])
	col.B = parseToUint8(segments[2])
	return
}
