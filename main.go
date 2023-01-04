package main

import (
	"fmt"
)

var PORT = 2022

func main() {
	// typeData()
	// typeArray()
	// typeSlice()
	// typeMap()
	// condisi()
	// hasil, angka := sayHellow("ebiebi")
	// fmt.Println(hasil)
	// fmt.Println(angka)
	// var data = []string{"wick", "john", "ethon"}
	// var hasil = filter(data, func(s string) bool {
	// 	return strings.Contains(s, "o")
	// })
	fmt.Println("hasil")
}

func typeData() {
	// var firstName bool = false
	var lastName = "ebi"
	hoby := "tidur"
	// var umur int = 12

	fmt.Printf("nama saya %s hoby %s \n", lastName, hoby)
	// fmt.Println("asas \n")
}

func typeArray() {
	var names [4]string
	names[0] = "nama1"
	names[1] = "nama2"
	// names[2] = "nama3"
	// names[3] = "nama4"

	var buah = [3]string{"apple", "mangga"}
	fmt.Println("jumlah capsity", cap(buah))
	fmt.Println("jumlah elemnt", len(buah))
	fmt.Println("isi elemnt", buah)

	var angka = [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println("jumlah capsity", cap(angka))
	fmt.Println("jumlah elemnt", len(angka))
	fmt.Println("isi elemnt", angka)

	// for i := 0; i < len(angka); i++ {
	// 	fmt.Println(angka[i])
	// }

	for index, value := range angka {
		fmt.Printf("index %d value %d \n", index, value)
	}
}

func typeSlice() {
	var buah = []string{"apple", "mangga", "rambutan", "durian"}
	var buaha = buah[0:3]
	var buahb = buaha[0:1]

	fmt.Println(buah)
	fmt.Println(buaha)
	fmt.Println(buahb)

	buahb[0] = "anggur"

	fmt.Println(buah)
	fmt.Println(buaha)
	fmt.Println(buahb)

}

func typeMap() {
	// var bulan = map[string]string{
	// 	"januari":  "satu",
	// 	"februari": "dua",
	// }

	var bulana = map[string]interface{}{
		"januari":  "satu",
		"februari": "dua",
	}

	var value, isExist = bulana["januari"]

	delete(bulana, "februari")

	fmt.Println(bulana)
	fmt.Printf("value %s exist %v", value, isExist)
}

func condisi() {
	var tahun = 2015

	if tahun > 2010 {
		println("true")

	} else if tahun < 2010 {
		println("true")

	} else {
		fmt.Println("false")
	}

	if lahir := 2021 - tahun; lahir > 20 {
		fmt.Println("anda sudah dewasa")
	} else if lahir < 10 {
		fmt.Println("anda masih bocah")
	} else {
		fmt.Println("anda remaja")
	}
}

func loopings() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	var start = 0

	for start < 10 {
		fmt.Println(start)
		start++
	}

	for {

		start++
		if start == 10 {
			break
		}
	}

	// for index, value := range angka {
	// 	fmt.Printf("index %d value %d \n", index, value)
	// }
}

func sayHellow(name string) (string, int) {
	return "helow", 20
}

func filter(data []string, callback func(string) bool) []string {
	var result []string

	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}

	return result
}
