package main

import "fmt"

func main() {
	//contoh variabel menggunakan tipe data
	var a string
	a = "Nizam"
	fmt.Println(a)

	a = "Deflorio"
	fmt.Println(a)

	var b int8
	b = 10
	fmt.Println(b)

	var c bool
	c = false
	fmt.Println(c)

	var d float32
	d = 1.5
	fmt.Println(d)

	//contoh variabel tidak menggunakan tipe data / langsung
	var name = "Deflorio Arya Nizam"
	fmt.Println("Nama saya :", name)

	//contoh variabel tidak menggunakan kata kunci var
	e := 1000
	fmt.Println(e)
}

// di Go-Lang membuat variabel bisa menggunakan dengan tipedata nya atau tidak

/*
di Go-Lang saat membuat variabel tidak wajib menggunakan var
Asalkan saat membuat variabel langsung menginialisasi datanya dengan menggunakan kata kunci :=
*/
