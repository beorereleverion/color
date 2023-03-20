package color

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
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

func RandomColor() Color {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	r := uint8(rnd.Intn(256))
	g := uint8(rnd.Intn(256))
	b := uint8(rnd.Intn(256))
	return RGBToColor(RGB{r, g, b})
}

func ParseHexColor(hexStr string) (Color, error) {
    hexStr = strings.TrimPrefix(hexStr, "#")

    if len(hexStr) != 6 {
        return 0, fmt.Errorf("invalid hex color format: %s", hexStr)
    }

    rgb := RGB{}

    r, err := strconv.ParseUint(hexStr[0:2], 16, 8)
    if err != nil {
        return 0, fmt.Errorf("invalid red component: %s", hexStr[0:2])
    }
    rgb.Red = uint8(r)

    g, err := strconv.ParseUint(hexStr[2:4], 16, 8)
    if err != nil {
        return 0, fmt.Errorf("invalid green component: %s", hexStr[2:4])
    }
    rgb.Green = uint8(g)

    b, err := strconv.ParseUint(hexStr[4:6], 16, 8)
    if err != nil {
        return 0, fmt.Errorf("invalid blue component: %s", hexStr[4:6])
    }
    rgb.Blue = uint8(b)

    return RGBToColor(rgb), nil
}