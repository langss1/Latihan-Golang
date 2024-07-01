package main

import "fmt"

func main() {
	var a int32 = 230
	var x int64 = int64(a)
	var z int8 = int8(a)

	fmt.Println(a)
	fmt.Println(x)
	fmt.Println(z)

	fmt.Println("")

	fmt.Println("------------------------------------------ ")

	fmt.Println("")

	var name = "Gilang Wasis"
	var e = name[0]
	var estring = string(e)

	fmt.Println(name)
	fmt.Println(estring)

}
