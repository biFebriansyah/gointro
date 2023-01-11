package module

import "fmt"

func ExamPointer() {
	var tanggal int = 20
	var pTanggal *int = &tanggal

	fmt.Println("Alamat 1", tanggal)
	fmt.Println("Alamat 2", *pTanggal)

	// *pTanggal = 80

	// fmt.Println("Alamat 1", tanggal)
	// fmt.Println("Alamat 2", *pTanggal)

	// var nama1 string = "hellow"
	// var nama2 string = nama1

	// fmt.Println("nama1", nama1)
	// fmt.Println("nama2", nama2)

	// nama2 = "haiii"

	// fmt.Println("nama1", nama1)
	// fmt.Println("nama2", nama2)
}

func echoName(name *string) *string {
	return name
}
