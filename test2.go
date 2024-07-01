package main

import "fmt"

func main() {
	var i, n, jumlah int64
	n = -300
	i = -3
	for i >= n {
		jumlah += i
		i = i - 3
	}
	fmt.Println(jumlah)
}
