package tempconv

import (
	"flag"
	"fmt"
)

// Celsius is float64 type.
type Celsius float64

// Fahrenheit is float64 type.
type Fahrenheit float64

type celsiusFlag struct{ Celsius }

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

// FToC converts temperature fro Fahrenheit to Celsius.
func FToC(f Fahrenheit) Celsius {
	c := (float64(f) - 32) * 5 / 9
	return Celsius(c)
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("Invalid temperature %q", s)
}

// CelsiusFlag defines a Celsius flag
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
