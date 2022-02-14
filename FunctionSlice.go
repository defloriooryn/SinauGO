package main

import "fmt"

func main() {

	bulan := [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	slice2 := bulan[2:10]

	//function len
	fmt.Println(len(slice2))

	//function cap
	fmt.Println(cap(slice2))

	// function append
	slice3 := append(slice2, "test")
	slice3[0] = "tes2"
	fmt.Println(slice3)
	fmt.Println(slice2)

	//function make
	newSlice := make([]uint, 5, 10)
	newSlice[0] = 1
	newSlice[2] = 2

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	//function copy
	copySlice := make([]uint, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

}

/*
Function Slice

- len(slice) = untuk mendapatkan panjang slice
- cap(slice) = untuk mendapatkan kapasitas
- append(slice, data) = membuat slice baru dengan menambah data ke posisi terakhir slice, jika kapasitas sudah penuh,
					    maka akan membuat array baru
- make([]TypeData, length, capacity) = membuat slice baru
- copy(destination, source) = menyalin slice dari source ke destination
*/
