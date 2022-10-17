package main

import "fmt"

func main() {
	for {
		fmt.Println("break for")
		break
	}

	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}
}
