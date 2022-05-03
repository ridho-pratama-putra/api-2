package utils

import (
	"fmt"
)

const PI float64 = 3.14 // Constant

func init() {
	fmt.Println("initializing two dimension")
}

func AreaOfCircleWithRadius(radius float64) float64{
	return PI * radius * radius
}
