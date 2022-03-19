package main

import (
	"fmt"
	"strconv"
)

// ReverseRunes returns its argument string reversed rune-wise left to right.
func ReverseRunes(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	// fmt.Println(ReverseRunes("makan"))
	var helloWorld string
	helloWorld = "hello world"
	fmt.Println("Reverse runes : ", ReverseRunes("makan"), " aslinya makan")
	fmt.Println("Makan 1 hari : ", 10, "x")
	fmt.Println("String hello world first initialization : ", helloWorld)
	helloWorld = "good morning"
	fmt.Println("String hello world after manipulation : ", helloWorld)

	fmt.Println("\nSIMPLE MATH FUNCTION------------------")
	eggsCounted := 10
	fmt.Println("int eggs counted direct init : ", eggsCounted)

	eggsCounted += 5
	fmt.Println("int eggs added by 5 : ", eggsCounted)

	eggsCounted -= 5
	fmt.Println("int diff d by 5 : ", eggsCounted)

	eggsCounted /= 4
	fmt.Println("int eggs divided by 4 : ", eggsCounted)

	fmt.Println("\nMULTIMPLE TIMES DEFINING------------------")
	var (
		firstname = "ta"
		lastname  = "kari"
	)
	fmt.Println("multiple times defining : ", firstname, lastname)

	fmt.Println("\nDEFINE ARRAY OF STRING------------------")
	var nameArray [3]string
	nameArray[0] = "Ju"
	nameArray[1] = "tak"
	fmt.Println(nameArray)
	fmt.Println(len(nameArray))
	fmt.Println(cap(nameArray))
	if nameArray[2] == "" {
		fmt.Println("second element array of string is empty")
	}

	fmt.Println("\nDEFINE ARRAY OF INT------------------")
	var nameArrayInt [3]int
	nameArrayInt[0] = 1
	nameArrayInt[1] = 2
	fmt.Println(nameArrayInt)
	fmt.Println(len(nameArrayInt))
	fmt.Println(cap(nameArrayInt))
	if nameArrayInt[2] == 0 {
		fmt.Println("second element array of int is empty")
	} else {
		fmt.Println("second element array of int is not empty")
	}

	if len(nameArrayInt) > 1 {
		fmt.Println("array int having len more than 1")
	}

	fmt.Println("\nMAKE STRING WITH LEN 2 CAP 4, it will increase capacity 2x when it was overflow due to append")
	newSlice := make([]string, 2, 4)
	newSlice[0] = "balqis"
	newSlice[1] = "balqis"
	newSlice = append(newSlice, "ayo")
	newSlice = append(newSlice, "ayo")
	newSlice = append(newSlice, "ayo")
	newSlice = append(newSlice, "ayo")
	newSlice = append(newSlice, "ayo")
	newSlice = append(newSlice, "ayo")
	newSlice = append(newSlice, "ayo")

	fmt.Println("LEN : ", len(newSlice))
	fmt.Println(cap(newSlice))
	for i, s := range newSlice {
		fmt.Println(i, "th element value: ", s)
	}

	fmt.Println("\nMAKE STRING WITH LEN 0 will generate dynamic array------------------")
	newSlice1 := make([]string, 0)
	fmt.Println("awal kapasitas array ", cap(newSlice1))
	newSlice1 = append(newSlice1, "balqis")
	fmt.Println("setelah pengisian awal kapasitas array ", cap(newSlice1))
	newSlice1 = append(newSlice1, "ayo")
	newSlice1 = append(newSlice1, "iya")
	newSlice1 = append(newSlice1, "lets go")
	newSlice1 = append(newSlice1, "lets gi")
	newSlice1 = append(newSlice1, "lets gl")
	for i := 0; i < 10; i++ {
		newSlice1 = append(newSlice1, `lets `+strconv.Itoa(i))
	}

	fmt.Println("panjang array ", len(newSlice1))
	fmt.Println("kapasitas array ", cap(newSlice1))
	for i := 0; i < len(newSlice1); i++ {
		fmt.Println(i, "th element value:", newSlice1[i])
	}

	fmt.Println("\nMAKE []INT WITH LEN 0------------------")
	s := make([]int, 0, 3)
	s = append(s, 0)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	fmt.Println("\nCREATE SLICE FROM ARRAY WITH := ")
	month := [...]string{ //create array
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	fullMonthCopy := month[0:] // create fullmonth copy as a slice
	fmt.Println(fullMonthCopy)

	fmt.Println("\nCOPYING SLICE AND REPLACE. it will affecting both array")
	month2 := fullMonthCopy[1:]
	fmt.Println("month2", month2)
	month2[1] = "alanDE`"            // affecting month2, fullMonthCopy, and month
	month2 = append(month2, "maret") // but append only affecting the destination array
	fmt.Println("month", month)
	fmt.Println("fullMonthCopy", fullMonthCopy)
	fmt.Println("month2", month2)

	fmt.Println("\nCOPYING SLICE WITH COPY FUNCTION")
	source := make([]string, 2, 2)
	source[0] = "senin"
	source[1] = "selasa"
	destination := make([]string, len(source), cap(source))
	copy(destination, source)
	fmt.Println(destination)

	// we cant append array
	//month = append(month, "wowNewMoon")
	//fmt.Println("\ncan we appendd array?")
	//fmt.Println(month)

	fmt.Println("\nCreate map")
	maps := make(map[string]string)
	maps["type"] = "LAPTOP"
	maps["price"] = "30000"
	maps["WILL BE DELETED SOON"] = "30000"
	fmt.Println("maps type ", maps["type"])
	delete(maps, "WILL BE DELETED SOON")
	fmt.Println("maps len ", len(maps))
	maps1 := map[string]string{
		"name":    "LENOVO",
		"factory": "hongkong",
	}
	fmt.Println("maps with direct init ", maps1)



	fmt.Println("\nCreate switch case")
	value := "m"
	switch value {
	case "mobile":
		fmt.Println("tilutlit")
	case "motor":
		fmt.Println("tringting")
	default:
		fmt.Println("default")
	}

	switch  {
	case len(value) >= 1:
		fmt.Println("len more than 1")
	case len(value) == 1:
		fmt.Println("len is 1")
	default:
		fmt.Println("default case")
	}


	fmt.Println("\nCreate for loop accessing array")
	accessArray := [...]string{
		"demons", "for", "you",
	}
	accessArray[2] = "uncle stan"
	for index, name := range accessArray {
		fmt.Println("val index: ", index, " value: ", name)
	}


	fmt.Println("\nCreate for loop accessing map")
	accessMap := map[string]string{
		"inject": "demons",
		"predicate":"for",
		"subject": "you",
	}
	accessMap["horn"] = "with horn"
	for index, name := range accessMap {
		fmt.Println("val index: ", index, " value: ", name)
	}


	fmt.Println("\nCreate function with parameter")
	tokoh("jalaludin rumi")
	tokoh("syeh bahaudin")
	tokohWithAge("syeh bahaudin", 23)


	fmt.Println("\nCreate function with single return value")
	if funcWithSingleReturnValue("merenung") == true {
		fmt.Println("lagi berdzikir")
	} else {
		fmt.Println("lagi tidur")
	}

	fmt.Println("\nCreate function with multiple return value")
	perluMinumObat, mandiBerapaKali := funcWithMultipleReturnValue(true)
	fmt.Println(perluMinumObat)
	fmt.Println(mandiBerapaKali)

	fmt.Println("\nCreate function ignoring some multiple return value")
	perluMinumObatKah, _ := funcWithMultipleReturnValue(true)
	fmt.Println(perluMinumObatKah)

	fmt.Println("\nCreate function with decided return aliasing ")
	jumlahSiswa, aa := funcWithNamedReturnValue()
	fmt.Println("sebanyak ", jumlahSiswa, " siswa sedang berwisata", aa)


	fmt.Println("\nCreate function with varArgs")
	totalBiaya, apakahPakaiKresek := funcWithVarArgsParameter(2, "pecel")
	fmt.Println("harus bayar ", totalBiaya, " ", apakahPakaiKresek)


	fmt.Println("\nCreate Variadic function varArgs value as slice")
	ints := []int{
		2, 4, 6,
	}
	fmt.Println("total element slice ", funcVariadicWithVarArgsSliceParameter(ints...))
}

func funcWithSingleReturnValue(reason string) bool {
	if reason == "merenung" {
		return true
	} else {
		return false
	}
}

func funcVariadicWithVarArgsSliceParameter(numbers ...int) int{
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func funcWithVarArgsParameter(quantityKG int, kind ...string) (rupiah int, pakaiKresek string) {
	hargaKiloan := 1000
	rupiah = hargaKiloan * quantityKG
	if len(kind) > 2 {
		pakaiKresek = "pakai kresek aja"
	} else {
		pakaiKresek = "gak pakai kresek"
	}
	return
}

func tokoh(name string){
	fmt.Println("tokoh : ", name)
}

func tokohWithAge(name string, age int){ // golang cannot overdrive
	fmt.Println("tokoh : ", name, " age: ", age)
}


func funcWithMultipleReturnValue(isHealthy bool) (string, string) {
	if isHealthy == true {
		return "tidak perlu minum obat", " mandi 5 kali"
	} else {
		return "perlu minum obat", " mandi 2 kali"
	}
}

func funcWithNamedReturnValue() (jumlahSiswa int, tujuanWisata string) {
	jumlahSiswa = 10
	tujuanWisata = " ke WC"
	return
}
