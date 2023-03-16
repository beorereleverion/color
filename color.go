package color

import (
	"fmt"
	"strconv"
	"strings"
)

type Color int

type RGB struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

func (c Color) HexString() string {
	return fmt.Sprintf("%X", c)
}

func (c Color) Int() int {
	return int(c)
}

func (c Color) GetRGB() RGB {
	r := uint8((c >> 16) & 0xFF)
	g := uint8((c >> 8) & 0xFF)
	b := uint8(c & 0xFF)
	return RGB{r, g, b}
}

func (r RGB) String() string {
	return fmt.Sprintf("%d,%d,%d", r.Red, r.Green, r.Blue)
}

func RGBToColor(rgb RGB) Color {
	color := int(rgb.Red)<<16 | int(rgb.Green)<<8 | int(rgb.Blue)
	return Color(color)
}

func ParseRGB(rgbStr string) (Color, error) {
	parts := strings.Split(strings.Trim(rgbStr, " "), ",")
	if len(parts) != 3 {
		return 0, fmt.Errorf("invalid RGB string format: %s", rgbStr)
	}

	r, err := strconv.ParseUint(strings.Trim(parts[0], " "), 10, 8)
	if err != nil {
		return 0, fmt.Errorf("invalid red component: %s", parts[0])
	}
	g, err := strconv.ParseUint(strings.Trim(parts[1], " "), 10, 8)
	if err != nil {
		return 0, fmt.Errorf("invalid green component: %s", parts[1])
	}
	b, err := strconv.ParseUint(strings.Trim(parts[2], " "), 10, 8)
	if err != nil {
		return 0, fmt.Errorf("invalid blue component: %s", parts[2])
	}

	color := RGBToColor(RGB{Red: uint8(r), Green: uint8(g), Blue: uint8(b)})
	return color, nil
}
