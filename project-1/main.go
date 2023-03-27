package main

import (
	"fmt"
	"sync"
)

func main() {

	var data1 = []string{
		"bisa 1",
		"bisa 2",
		"bisa 3",
	}

	var data2 = []string{
		"coba 1",
		"coba 2",
		"coba 3",
	}

	fmt.Println("==============ACAK==============")

	var wg sync.WaitGroup
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go func(id int) {
				fmt.Printf("%v %d\n", data1, id)
			wg.Done()
		}(i)
		wg.Add(1)
		go func(id int) {
				fmt.Printf("%v %d\n", data2, id)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("==============ACAK==============")
}

// blm selese msh acak aja tpi dikumpulkan dlu