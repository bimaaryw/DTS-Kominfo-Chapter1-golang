package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	var data1 interface{} = []string{
		"bisa 1",
		"bisa 2",
		"bisa 3",
	}

	var data2 interface{} = []string{
		"coba 1",
		"coba 2",
		"coba 3",
	}

	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	for i := 1; i <= 8; i++ {
		wg.Add(1)
		go func(id int) {
			if id%2 == 1 {
				fmt.Printf("%v %d\n", data1, id)
			} else {
				fmt.Printf("%v %d\n", data2, id)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// blm selese msh proses tpi sya kumpulkan dlu