package main

import (
	"fmt"
)

func main() {
	type NoKTP string
	type married bool

	var noktpeko NoKTP = "1234566"
	var marriedstatus married = true
	fmt.Println(noktpeko)
	fmt.Println(marriedstatus)

}
