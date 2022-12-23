package extendedcolor

import (
	"fmt"
	"image/color"
)

func HexFromRgbColor(color color.RGBA) string {
	return fmt.Sprintf("#%X%X%X", color.R, color.G, color.B)
}

func ParseHexColor(s string) (c color.RGBA, err error) {
	c.A = 0xff
	switch len(s) {
	case 7:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 4:
		_, err = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		// Double the hex digits:
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		err = fmt.Errorf("invalid length, must be 7 or 4")

	}
	return
}

func MapManyHex(inputs []string) []color.RGBA {
	var items []color.RGBA
	for _, item := range inputs {
		col, err := ParseHexColor(item)
		if err == nil {
			items = append(items, col)
		}
	}

	return items
}