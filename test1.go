package main

import (
	"fmt"
)

func main() {
	var total, b, i int64
	b = 5
	i = 2
	total = 0
	for i <= 1000000002 {
		total = total + b
		i = i + b
	}
	fmt.Println(total)

}
