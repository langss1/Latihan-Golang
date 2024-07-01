package main

import (
	"fmt"
)

func main() {
	var total, b int64
	b = 2
	for b <= 1000000003 {
		total = total + b
		b = b + 7
	}
	fmt.Println(total)
}
