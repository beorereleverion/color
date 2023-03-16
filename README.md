# Color Library

This library provides predefined colors and utility functions for working with colors.

## Usage

To use the library, simply import it and access the predefined color variables. For example, to use the `Green` color:

```go
package main

import (
    "fmt"

    "github.com/beorereleverion/color"
)

func main() {
    rgb := color.Green.GetRGB()
    fmt.Printf("\033[38;2;%d;%d;%dmGreen line\033[0m\n", rgb.Red, rgb.Green, rgb.Blue)
    fmt.Printf("Green with hex %s has next rgb: %v and int %d", color.Green.HexString(), rgb, color.Green)
}
```

In this example, we use the Green color from the color package and print a green line to the console using the RGB values of the color. We also print the RGB values and hex string of the color using the GetRGB and HexString functions provided by the Green color variable.
Predefined Colors

## License

This library is licensed under the MIT License. See the LICENSE file for details.

## Contributing

Contributions are welcome! Please submit a pull request or file an issue on the repository.