package main

import (
	"./pkg_string"
	"fmt"
)

func main() {
	fmt.Printf(stringutil.Reverse("!oG , olleH") + "\n")
	//var tempstr string
	tempstr := "123"
	//fmt.Scanf("%s", &tempstr)
	fmt.Scan(&tempstr)
	fmt.Printf(tempstr + "\n")
	fmt.Print(&tempstr)
}
