package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BolingC       Celsius = 100
)

func (c Celsius) String() {
	return fmt.Sprintf("%g℃", c)
}

func (c Fahrenheit) String() {
	return fmt.Sprintf("%g℉", f)
}
