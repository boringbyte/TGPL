package tempconv

// CToF converts Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts Fahrenheit to Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CToK converts Celsius to Kelvin
func CToK(c Celsius) Kelvin { return Kelvin(c - 273.15) }

// KToC converts Kelvin to Celsius
func KToC(k Kelvin) Celsius { return Celsius(k + 273.15) }
