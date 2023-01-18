package basic

import "fmt"

func ExampDeffer() {
	db := "connetction"
	queryData(db)
	query2(db)
}

func queryData(db string) {
	defer fmt.Println("tutup query data")
	fmt.Println("run query to db")
	fmt.Println("run query2 to db")

}

func query2(db string) {
	defer fmt.Println("tutup query2")
	fmt.Println("run query3 to db")
}
