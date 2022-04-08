package main

import (
	_ "api-2/example/stageThree/unusedUtils"
	"api-2/example/stageThree/utils"
	"fmt"
	"os"
)

func main() {
	fmt.Println("/nPACKAGE & IMPORTS")
	fmt.Println("stage three", utils.AlayString("Ini nih "))
	fmt.Println("stage three", utils.AreOfCircleWithRadius(5.1))

	fmt.Println("/nACCESS MODIFIER")
	fmt.Println("accessing value defined in twoDimension", utils.PI)

	fmt.Println("/nPACKAGE OS")
	// go run example/stageThree/stageThree.go arema indonesia
	fmt.Println("args 0, command itu sendiri :", os.Args[0]) //  /var/folders/ng/4wp9dvzn65lf1nbvhc38wt0cnwlks5/T/go-build1796574156/b001/exe/stageThree
	//fmt.Println("apa sih ini :", os.Args[1]) // arema
	name, err := os.Hostname()
	fmt.Println("host name:", name)
	fmt.Println("current PATH variable system:", os.Getenv("PATH"))

	if err == nil {
		fmt.Println("no error")
	} else if err != nil {
		fmt.Println("error", err)
	}



}
