package main

import "fmt"

func main() {
	var name = "eko"
	var e byte = name[1]
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(eString)

}
