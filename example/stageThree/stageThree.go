package main

import (
	_ "api-2/example/stageThree/unusedUtils"
	"api-2/example/stageThree/utils"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
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

	fmt.Println("\nPACKAGE STRCONV")
	seribu64, err := strconv.ParseInt("999999999999999999999", 10, 64)
	seribu32, err := strconv.ParseInt("999999999999999999999", 10, 32)
	seribu16, err := strconv.ParseInt("999999999999999999999", 10, 16)
	seribu8, err := strconv.ParseInt("999999999999999999999", 10, 8)
	fmt.Println(seribu64)
	fmt.Println(seribu32)
	fmt.Println(seribu16)
	fmt.Println(seribu8)
	ji, _:= strconv.ParseInt("-354634382", 10, 32)
	fmt.Printf("this string %v has become pure %T \n", ji, ji)
	k := strconv.FormatBool(false)
	fmt.Println("this bool %s has become string ", k)
	j, _:= strconv.ParseBool("trueee")
	fmt.Printf("this string bool %t has become pure bool \n", j)
	kf := strconv.FormatFloat(1.8, 'E', -1, 32)
	fmt.Printf("this float %v has become string \n", kf)
	jf, _:= strconv.ParseFloat("1.8", 32)
	fmt.Printf("this string float %b has become pure float ", jf)

	fmt.Println("\nPACKAGE MATH")
	var floatValue float64 = 1.50;
	var floatComparison float64 = 1.3
	fmt.Printf("origin value %v will rounded as %v\n",floatValue,math.Round(floatValue))
	fmt.Printf("origin value %v will floor as %v\n",floatValue,math.Floor(floatValue))
	fmt.Printf("origin value %v will ceil as %v\n",floatValue,math.Ceil(floatValue))
	fmt.Printf("between %v and %v, value %v is bigger\n",floatValue, floatComparison, math.Max(floatValue, floatComparison))
	fmt.Printf("between %v and %v, value %v is smaller\n",floatValue, floatComparison, math.Min(floatValue, floatComparison))

}
