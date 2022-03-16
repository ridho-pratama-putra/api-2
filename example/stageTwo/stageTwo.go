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

func main() {
	fmt.Println("\nFunction value, like aliasing function")
	apakahLagiMerenung := funcWithSingleReturnValue

	result := apakahLagiMerenung("merenung")
	fmt.Println(" function aliassing result : ", result)

	fmt.Println("\nFunction as parameter")
	functWithFuncAsParameter("merenung", funcWithSingleReturnValue)
	functWithFuncAsParameter("lari-lari", funcWithSingleReturnValue)

}
