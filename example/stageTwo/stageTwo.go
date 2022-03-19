package main

import (
	"errors"
	"fmt"
)

func funcWithSingleReturnValue(reason string) bool {
	if reason == "merenung" {
		return true
	} else {
		return false
	}
}

func functWithFuncAsParameter(ngapain string, funcAsParameter func(string) bool) {
	result := funcAsParameter(ngapain)
	if result == true {
		fmt.Println(" example string param value: ", ngapain, ". example result true from func as param: ", result)
	} else {
		fmt.Println(" example string param value: ", ngapain, ". example result false from func as param: ", result)
	}
}

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) == true {
		fmt.Println(" You're blocked ", name)
	} else {
		fmt.Println(" welcome ", name)
	}
}

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	}
	return value * factorialRecursive(value-1)
}

func logging() {
	fmt.Println(" selesai memanggil function")
}

func loggingLast() {
	fmt.Println(" selesai memanggil function LAST")
}

func runApplication(value int) {
	defer logging()
	defer loggingLast()
	result := 10 / value
	fmt.Println(" function result", result)
	fmt.Println(" run application")
}

func main() {
	fmt.Println("\nFunction value, like aliasing function")
	apakahLagiMerenung := funcWithSingleReturnValue

	result := apakahLagiMerenung("merenung")
	fmt.Println(" function aliassing result : ", result)

	fmt.Println("\nFunction as parameter")
	functWithFuncAsParameter("merenung", funcWithSingleReturnValue)
	functWithFuncAsParameter("lari-lari", funcWithSingleReturnValue)

	fmt.Println("\nFunction anonymous, direct function in param or function without name")
	blacklist := func(name string) bool {
		return name == "amrozi"
	}
	registerUser("amrozi", blacklist) // by variable
	registerUser("anjing", func(name string) bool { // by parameter
		return name == "amrozi"
	})

	fmt.Println("\nFunction recursive fatorial loop")
	fmt.Println(" result ", factorialLoop(4))
	fmt.Println(" result recursive ", factorialRecursive(4))

	fmt.Println("\nFunction closure")
	name := "gedang"
	counter := 0
	increment := func() {
		name := "telo"
		fmt.Println(" do Increment")
		fmt.Println(" inside closure", name)
		counter++
	}

	increment()
	fmt.Println(" counter value", counter)
	fmt.Println(" outside of closure", name)

	fmt.Println("\nFunction DEFER. executed (LIFO if multiple) di akhir bahkan jika ada error.")
	runApplication(10)
	fmt.Println("\nFunction DEFER with PANIC untuk break program")
	//runApplication(0) // division by zero will return panic and break following lines
	fmt.Println("\nFunction RECOVER. better place on defer")
	isPanic(true)

	fmt.Println("\nSTRUCT. as representasi dari data. mirip class di java")
	var people Customer
	people.Address = "jalan gajayana"
	people.Name = "sam kodir"
	people.Age = 25
	fmt.Println("here is customer", people)

	fmt.Println("\nSTRUCT literals #1. direct initialize")
	jordan := Customer{
		Name:    "jordan",
		Address: "america",
		Age:     25,
	}
	fmt.Println(jordan)

	fmt.Println("\nSTRUCT literals #2. direct initialize. harus urutl. AllArgsConstructor")
	budi := Customer{"biji", "bunga", 99}
	fmt.Println(budi)

	fmt.Println("\nSTRUCT methods")
	rubik := Customer{"rubik", "jalan", 34}
	rubik.sayHello()

	fmt.Println("\nINTERFACE. blue print")
	gunawan := Customer{"gunawan", "jalan", 34}
	SayYourIdentity(gunawan)
	yanto := Seller{Name: "yanto", StoreName: "pakaian bayi"}
	SayYourIdentity(yanto)

	fmt.Println("\nNil untuk pengecekan null, hanya berlaku untuk interface, function, map ,slice, pointer dan channel")

	fmt.Println("\nError")
	contohError := errors.New(" contoh error")
	fmt.Println(contohError.Error())
	int, _ := pembagian(6, 3)
	fmt.Println(" no error, result", int)
	_, error := pembagian(6, 0)
	fmt.Println(" ada error", error)

	fmt.Println("\nType Assertion. replace data type ketika menggunakan interface kosong")
	resultTypeAssertion := random()
	stringConverted := resultTypeAssertion.(string)
	fmt.Println(" string converted", stringConverted)

	fmt.Println("\nType Assertion using switch case")
	resultUsingSwitch := randomBool()
	switch value := resultUsingSwitch.(type) {
	case string:
		fmt.Println(" result is string", value)
	case bool:
		fmt.Println(" result is bool", value)
	case int8:
		fmt.Println(" result is int", value)
	default:
		fmt.Println(" result is not string/int", value)
	}
}
func randomBool() interface{} {
	return "sad"
}
func random() interface{} {
	return "1"
}

func pembagian(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("pembagian dengan 0")
	}
	return a / b, nil
}

type Customer struct {
	Name, Address string
	Age           int
}
type Seller struct {
	Name, Address, StoreName string
	Age                      int
}

func (customer Customer) sayHello() {
	fmt.Println(" hallo", customer.Name)
}

type HasDetail interface {
	// can be used for another struct too ass
	GetName() string
	GetAddress() string
	GetAuthority() string
}

func (customer Customer) GetName() string { // contract for Customer
	return customer.Name
}

func (customer Customer) GetAddress() string { // contract for Customer
	return customer.Address
}

func (seller Seller) GetName() string { // contract for Seller
	return seller.Name
}

func (seller Seller) GetAddress() string { // contract for Seller
	return seller.Address
}

func (seller Seller) GetAuthority() string {
	return "Seller"
}

func (customer Customer) GetAuthority() string {
	return "Customer"
}

func SayYourIdentity(detail HasDetail) {
	fmt.Println(" hello my name is", detail.GetName())
	fmt.Println(" I live in", detail.GetAddress())
	fmt.Println(" I here as", detail.GetAuthority())
}

func isPanic(error bool) {
	defer loggingWithRecover()
	defer loggingWithRecoverLast()
	if error {
		panic("ERROR HAPPEN ")
	}
}

func loggingWithRecover() {
	errorMessage := recover()
	if errorMessage != nil {
		fmt.Println("error", errorMessage)
	} else {
		fmt.Println("no error")
	}
	fmt.Println("done executing")
}

func loggingWithRecoverLast() {
	errorMessage := recover()
	if errorMessage != nil {
		fmt.Println("error LAST", errorMessage)
	} else {
		fmt.Println("no error LAST")
	}
	fmt.Println("done executing LAST")
}
