package color

import "testing"

func TestGetRGB(t *testing.T) {
	c := Color(0xFFA500) // Orange
	expected := RGB{255, 165, 0}
	actual := c.GetRGB()

	if actual != expected {
		t.Errorf("Expected %+v, but got %+v", expected, actual)
	}
}

func TestHexString(t *testing.T) {
	color := Color(0xFFA500) // Orange
	expected := "FFA500"
	actual := color.HexString()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}

func TestRGBToColor(t *testing.T) {
	rgb := RGB{255, 165, 0} // Orange
	expected := Color(0xFFA500)
	actual := RGBToColor(rgb)

	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func TestParseRGB(t *testing.T) {
	// Test a valid RGB string
	rgbStr := "255, 165, 0"
	expected := Color(0xFFA500)
	actual, err := ParseRGB(rgbStr)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}

	// Test an invalid RGB string with too few components
	rgbStr = "255, 165"
	_, err = ParseRGB(rgbStr)
	if err == nil {
		t.Error("Expected an error, but got none")
	}

	// Test an invalid RGB string with too many components
	rgbStr = "255, 165, 0, 0"
	_, err = ParseRGB(rgbStr)
	if err == nil {
		t.Error("Expected an error, but got none")
	}

	// Test an invalid RGB string with an invalid component value
	rgbStr = "255, 165, 256"
	_, err = ParseRGB(rgbStr)
	if err == nil {
		t.Error("Expected an error, but got none")
	}
}

func TestRGBString(t *testing.T) {
	r := RGB{Red: 255, Green: 0, Blue: 128}
	expected := "255,0,128"
	result := r.String()
	if result != expected {
		t.Errorf("RGB.String() returned unexpected result: got %s, expected %s", result, expected)
	}
}

func TestColorInt(t *testing.T) {
	c := Color(0xff0000)
	if c.Int() != 16711680 {
		t.Errorf("Expected 16711680, but got %d", c.Int())
	}
}

func TestRandomColor(t *testing.T) {
	color := RandomColor()
	if color.Int() < 0 || color.Int() > 0xFFFFFF {
		t.Errorf("RandomColor() returned invalid color: %v", color)
	}
}

func TestParseHexColor(t *testing.T) {
    // Test valid hex color
    hex := "#FF0000"
    expected := RGB{Red: 255, Green: 0, Blue: 0}
    color, err := ParseHexColor(hex)
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
    if color != RGBToColor(expected) {
        t.Errorf("Expected color %v but got %v", expected, color)
    }

    // Test invalid hex color
    hex = "00FF"
    _, err = ParseHexColor(hex)
    if err == nil {
        t.Error("Expected error for invalid hex color but got nil")
    }
}