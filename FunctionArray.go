package main

import "fmt"

func main() {

	a := [3]byte{
		1,
		2,
		3,
	}

	fmt.Println(len(a))
	fmt.Println(a[0])
	a[2] = 4
	fmt.Println(a[2])

}

// len(array) : untuk mendapatkan panjang array
// array[index] : untuk mendapatkan data di posisi index
// array[index] = value : untuk merubah data di posisi index
