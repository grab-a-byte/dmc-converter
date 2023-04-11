package extendedcolor

import (
	"image/color"
	"strconv"
	"strings"
)

func Compare(col1 color.RGBA, col2 color.RGBA) int {
	var diff int = subtractHighestFromLowest(col1.R, col2.R) +
		subtractHighestFromLowest(col1.G, col2.G) +
		subtractHighestFromLowest(col1.B, col2.B)

	return diff
}

func MapCommaSeperatedLinesToColors(inputs []string) []color.RGBA {
	return mapMany(fromCommaSepeartedLine, inputs)
}

func MapSpaceSeperatedLinesToColors(inputs []string) []color.RGBA {
	return mapMany(fromSpaceSeparatedLine, inputs)
}

func subtractHighestFromLowest(i1, i2 uint8) int {
	if i1 > i2 {
		return int(i1 - i2)
	}
	return int(i2 - i1)
}

func mapMany(mapper func(string) color.RGBA, inputs []string) []color.RGBA {
	var colors []color.RGBA
	for _, item := range inputs {
		colors = append(colors, mapper(item))
	}
	return colors
}

func parseToUint8(str string) uint8 {
	value, _ := strconv.ParseUint(str, 10, 8)
	return uint8(value)
}

func fromSpaceSeparatedLine(str string) color.RGBA {
	return fromLine(" ", str)
}

func fromCommaSepeartedLine(str string) color.RGBA {
	return fromLine(",", str)
}

func fromLine(splitter string, input string) (col color.RGBA) {
	segments := strings.Split(input, splitter)
	col.R = parseToUint8(segments[0])
	col.G = parseToUint8(segments[1])
	col.B = parseToUint8(segments[2])
	return
}
