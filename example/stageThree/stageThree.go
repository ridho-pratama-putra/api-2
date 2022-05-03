package main

import (
	_ "api-2/example/stageThree/unusedUtils"
	"api-2/example/stageThree/utils"
	"container/list"
	"container/ring"
	"flag"
	"fmt"
	"math"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("\nPACKAGE & IMPORTS")
	fmt.Println("stage three", utils.AlayString("Ini nih "))
	fmt.Println("utils to calculate AreaOfCircleWithRadius", utils.AreaOfCircleWithRadius(5.1))

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
	var port *int = flag.Int("port", 8080, "please input number")
	// help clue will appear triggered with --help or error value supplied
	flag.Parse()
	fmt.Println(*host)
	fmt.Println(*port)

	fmt.Println("\nPACKAGE STRING")
	fmt.Println(strings.ToTitle("stringto title"))
	fmt.Println(strings.Trim(" trim P@ssw0rd ", " "))
	fmt.Println(strings.ToLower("THIS ALL CAPItaL"))
	fmt.Println(strings.ToUpper("THIS ALL small"))
	var stringElement []string = strings.Split("this will split", " ")
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
	ji, _ := strconv.ParseInt("-354634382", 10, 32)
	fmt.Printf("this string %v has become pure %T \n", ji, ji)
	k := strconv.FormatBool(false)
	fmt.Printf("this bool %s has become string ", k)
	j, _ := strconv.ParseBool("trueee")
	fmt.Printf("this string bool %t has become pure bool \n", j)
	kf := strconv.FormatFloat(1.8, 'E', -1, 32)
	fmt.Printf("this float %v has become string \n", kf)
	jf, _ := strconv.ParseFloat("1.8", 32)
	fmt.Printf("this string float %b has become pure float ", jf)

	fmt.Println("\nPACKAGE MATH")
	var floatValue float64 = 1.50;
	var floatComparison float64 = 1.3
	fmt.Printf("origin value %v will rounded as %v\n", floatValue, math.Round(floatValue))
	fmt.Printf("origin value %v will floor as %v\n", floatValue, math.Floor(floatValue))
	fmt.Printf("origin value %v will ceil as %v\n", floatValue, math.Ceil(floatValue))
	fmt.Printf("between %v and %v, value %v is bigger\n", floatValue, floatComparison, math.Max(floatValue, floatComparison))
	fmt.Printf("between %v and %v, value %v is smaller\n", floatValue, floatComparison, math.Min(floatValue, floatComparison))

	fmt.Println("\nPACKAGE CONTAINER/LIST")
	data := list.New()
	data.PushBack(3)
	data.PushBack(4)
	data.PushBack("lima")
	fmt.Print(data.Front().Value)
	fmt.Print(data.Back().Value)
	fmt.Println(data.Front().Next().Value)
	// forward
	for a := data.Front(); a != nil; a = a.Next() {
		if a.Next() == nil {
			fmt.Println(a.Value)
		} else {
			fmt.Print(a.Value)
		}
	}
	// backward
	for a := data.Back(); a != nil; a = a.Prev() {
		if a.Prev() == nil {
			fmt.Println(a.Value)
		} else {
			fmt.Print(a.Value)
		}
	}
	// try add some element
	data.PushFront("enol")
	for a := data.Front(); a != nil; a = a.Next() {
		if a.Next() == nil {
			fmt.Println(a.Value)
		} else {
			fmt.Print(a.Value)
		}
	}

	fmt.Println("\nPACKAGE CONTAINER/RING")
	var dataRing *ring.Ring = ring.New(5) // not dynamic like list
	for i := 0; i <= data.Len(); i++ {
		//fmt.Println(i)
		dataRing.Value = "Value " + strconv.FormatInt(int64(i), 10)
		dataRing = dataRing.Next()
	}

	//fmt.Println(*dataRing) // show memory location because kind of doesnt have println
	// {0xc0000c2000 0xc0000c2060 <nil>}

	dataRing.Do(func(value interface{}) {
		fmt.Println(value)
	})

	fmt.Println("\nPACKAGE sort")
	var users []User = []User{
		{Name: "A", Age: 10},
		{Name: "B", Age: 13},
		{Name: "C", Age: 12},
	}

	sort.Sort(UserSlice(users))
	fmt.Println(users)

	fmt.Println("\nPACKAGE time")
	//ANSIC			"Mon Jan _2 15:04:05 2006"
	//UnixDate		"Mon Jan _2 15:04:05 MST 2006"
	//RubyDate		"Mon Jan 02 15:04:05 -0700 2006"
	//RFC822		"02 Jan 06 15:04 MST"
	//RFC822Z		"02 Jan 06 15:04 -0700"
	//RFC850		"Monday, 02-Jan-06 15:04:05 MST"
	//RFC1123		"Mon, 02 Jan 2006 15:04:05 MST"
	//RFC1123Z		"Mon, 02 Jan 2006 15:04:05 -0700"
	//RFC3339		"2006-01-02T15:04:05Z07:00"
	//RFC3339Nano	"2006-01-02T15:04:05.999999999Z07:00"
	//Stamp			"3:04PM"
	//StampMilli	"Jan _2 15:04:05.000"
	//StampMicro	"Jan _2 15:04:05.000000"
	//StampNano		"Jan _2 15:04:05.000000000"
	location, _ := time.LoadLocation("Asia/Jakarta")
	fmt.Println("full", time.Now().In(location))
	fmt.Println("years", time.Now().Year())
	fmt.Println("month", time.Now().Month())
	fmt.Println("day", time.Now().Day())     // tanggal
	fmt.Println("day", time.Now().Weekday()) // hari
	fmt.Println("hour", time.Now().Hour())
	fmt.Println("minute", time.Now().Minute())
	fmt.Println("second", time.Now().Second())
	fmt.Println("create date UTC", time.Date(2009, time.November, 11, 19, 0, 0, 0, time.UTC))
	parsedTime, _ := time.Parse(time.RFC3339, "2006-01-31T15:04:55Z")
	fmt.Println("parse date UTC using RFC3339", parsedTime)

	rfc339DateFormat := "2006-01-02"
	parsedTimeCustom, _ := time.Parse(rfc339DateFormat, "2020-05-30")
	fmt.Println("parse date UTC using custom format", parsedTimeCustom)

	fmt.Println("\nPACKAGE reflect")
	womenUser := User{Name: "balqis", Age: 27, Address: "ordow"}
	noAddressUser := User{Name: "balqis", Age: 27}
	userType := reflect.TypeOf(womenUser)
	nameField := userType.Field(0)
	ageField := userType.Field(1)
	requiredNameField := nameField.Tag.Get("required")
	requiredAddress := userType.Field(2).Tag.Get("required")
	maxAddress := userType.Field(2).Tag.Get("max")
	fmt.Println("womenUser :: ", womenUser)
	fmt.Println("userType :: ", userType)
	fmt.Println("nameField :: ", nameField)
	fmt.Println("ageField :: ", ageField)
	fmt.Println("requiredName :: ", requiredNameField)                                           // return empty
	fmt.Println("requiredAddress :: ", reflect.TypeOf(requiredAddress), " :: ", requiredAddress) // return true
	fmt.Println("maxAddress :: ", reflect.TypeOf(maxAddress), " :: ", maxAddress)                // return true
	fmt.Println("type of Valueof:: ", reflect.TypeOf(reflect.ValueOf(maxAddress)))
	intConvertedMaxAddress, _ := strconv.ParseInt(maxAddress, 10, 8)
	fmt.Println("maxAddressVal :: ", intConvertedMaxAddress > 90)
	fmt.Println("isValid struct :: ", IsValid(womenUser))
	fmt.Println("isValid struct :: ", IsValid(noAddressUser))

	fmt.Println("\nPACKAGE regex")
	var regexpRule *regexp.Regexp = regexp.MustCompile("e[a-zA-Z]o")
	fmt.Println("applied rule for :: ", "eJo", " :: ", regexpRule.MatchString("eJo"))
	fmt.Println("applied rule for :: ", "e3o", " :: ", regexpRule.MatchString("e3o"))


}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i :=0; i<t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			return reflect.ValueOf(data).Field(i).Interface() != ""
		}
	}
	return true
}

type User struct {
	Name    string
	Age     int
	Address string `required:"true" max:"10"`
}

type UserSlice []User

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}
