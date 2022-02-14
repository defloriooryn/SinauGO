package main

import "fmt"

func main() {

	var values = [3]int8{1, 2, 3}
	fmt.Println(values)

	values_1 := [2]string{

		"Deflorio",
		"Arya",
	}

	fmt.Println(values_1)
}

//buat array tidak bisa pakai constant
//jika belum tahu berapa kapasitas aray nya maka pakai [...]
