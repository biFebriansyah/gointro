package intermediate

import (
	"fmt"
	"strconv"
	"sync"
)

var wg = sync.WaitGroup{}

func WaitGp() {
	fmt.Println("main running")

	wg.Add(2) // regist go routine

	go getChar("hello world")
	go getDigit([]int{1, 2, 3, 4, 5})

	wg.Wait()

	fmt.Println("main stop")
}

func getChar(st string) {
	var result string

	for _, c := range st {
		result += string(c)
	}

	fmt.Println(result)
	wg.Done()
}

func getDigit(dt []int) {
	var result string

	for _, c := range dt {
		result += strconv.Itoa(c)
	}

	fmt.Println(result)
	wg.Done()
}
