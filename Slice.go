package main

import "fmt"

func main() {

	hari := [...]string{
		"senin",
		"selasa",
		"rabu",
		"kamis",
		"jumat",
		"sabtu",
		"minggu",
	}

	var slice1 = hari[3:5]
	fmt.Println(slice1)
}

/*
membuat slice
array[low:high]
array[low:]
array[:high]
array[:]
*/
