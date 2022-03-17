package main

import (
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
		fmt.Println(" You're blocked ",name)
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
	return value * factorialRecursive(value - 1)
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
	fmt.Println(" function result",result)
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
	blacklist := func (name string) bool {
		return name == "amrozi"
	}
	registerUser("amrozi", blacklist) // by variable
	registerUser("anjing", func (name string) bool { // by parameter
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
	fmt.Println(" counter value",counter)
	fmt.Println(" outside of closure", name)


	fmt.Println("\nFunction DEFER. executed (LIFO if multiple) di akhir bahkan jika ada error.")
	runApplication(10)
	fmt.Println("\nFunction DEFER with PANIC untuk break program")
	//runApplication(0) // division by zero will return panic and break following lines
	fmt.Println("\nFunction RECOVER. better place on defer")
	isPanic(true)



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
