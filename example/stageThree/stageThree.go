package main

import (
	_ "api-2/example/stageThree/unusedUtils"
	"api-2/example/stageThree/utils"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("\nPACKAGE & IMPORTS")
	fmt.Println("stage three", utils.AlayString("Ini nih "))
	fmt.Println("stage three", utils.AreOfCircleWithRadius(5.1))

	fmt.Println("\nACCESS MODIFIER")
	fmt.Println("accessing value defined in twoDimension", utils.PI)

	fmt.Println("\nPACKAGE OS")
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

	fmt.Println("\nPACKAGE FLAG") // untuk parsing args command
	host := flag.String("host", "localhost", "this is help for clue")
	var port *int= flag.Int("port", 8080, "please input number")
	// help clue will appear triggered with --help or error value supplied
	flag.Parse()
	fmt.Println(*host)
	fmt.Println(*port)

	fmt.Println("\nPACKAGE STRING")
	fmt.Println(strings.ToTitle("stringto title"))
	fmt.Println(strings.Trim(" trim P@ssw0rd ", " "))
	fmt.Println(strings.ToLower("THIS ALL CAPItaL"))
	fmt.Println(strings.ToUpper("THIS ALL small"))
	var stringElement []string= strings.Split("this will split", " ")
	fmt.Println("RESULT", stringElement[0], stringElement[1], stringElement[2])
	fmt.Println("string is contain 1??", strings.Contains("who are you", "1"));
	fmt.Println("string replace all", strings.ReplaceAll("who are you", " ", "*"));
}
