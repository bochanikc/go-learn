package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a1, b1 int64
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &a1)
	in = bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &b1)

	fmt.Println(summOfDigits(a1, b1))
}
func summOfDigits(a, b int64) int64 {
	return a + b
}
