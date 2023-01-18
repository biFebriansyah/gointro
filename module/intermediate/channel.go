package intermediate

import "fmt"

func ChanEl() {

	wg.Add(2)
	var ch = make(chan int, 2) // bufer chanel

	go rOne(ch)
	go rTwo(ch)

	wg.Wait()

}

/*
Main \\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
Routen 1 \\\\\\\\\\\\\\\\\
Routen 2 \\\\\\\\\\\\\\\\\
*/

// get data
func rOne(ch chan int) {
	for data := range ch {
		fmt.Println("datanya = ", data)
	}

	wg.Done()
}

// routine send data
func rTwo(ch chan int) {
	defer func() {
		close(ch)
		wg.Done()
	}()

	ch <- 20 // write channel
	ch <- 50
}
