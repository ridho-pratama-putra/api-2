package main

import (
	_ "api-2/example/stageThree/unusedUtils"
	"api-2/example/stageThree/utils"
	"fmt"
)

func main() {
	fmt.Println("PACKAGE & IMPORTS")
	fmt.Println("stage three", utils.AlayString("Ini nih "))
	fmt.Println("stage three", utils.AreOfCircleWithRadius(5.1))

	fmt.Println("ACCESS MODIFIER")
	fmt.Println("accessing value defined in twoDimension", utils.PI)
}
