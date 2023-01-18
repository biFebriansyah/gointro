package basic

import "fmt"

type kurir struct {
	kota string
	prov string
}

func (k kurir) GetProv() string {
	return k.prov
}

type profile struct {
	email string
	phone int
}

type user struct {
	username string
	password string
	profile
}

func ExamStruct() {
	var data1 kurir
	data1.kota = "jakarta"
	data1.prov = "DKI"

	var us1 user
	us1.profile.email = "test@email.com"

	fmt.Println(us1)

}
