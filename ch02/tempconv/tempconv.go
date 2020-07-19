package tempconv

import "fmt"

// Celsius is a float64 type
type Celsius float64

// Fahrenheit is a float64 type
type Fahrenheit float64

// Constant of Celsius
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

// Kelvin is a float64 type
type Kelvin float64

// Constant of Kelvin
const (
	AbsoluteZeroK Kelvin = 0
	FreezingK     Kelvin = 273.15
	BoilingK      Kelvin = 373.15
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }
