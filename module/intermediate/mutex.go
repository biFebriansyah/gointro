package intermediate

import (
	"fmt"
	"sync"
)

var angka = 0
var mt = sync.Mutex{}
var wgs = sync.WaitGroup{}

func MuteEx() {

	for i := 0; i < 5; i++ {
		wgs.Add(2)
		mt.Lock()
		go showAngka()
		mt.Lock()
		go incement()
	}

	wgs.Wait()
}

func showAngka() {
	fmt.Printf("hello %d \n", angka)
	mt.Unlock()
	wgs.Done()
}

func incement() {
	angka++
	mt.Unlock()
	wgs.Done()
}
