package main

import "fmt"

func main(){
	var bulan = [...]string{
		"janauri",
		"februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"oktober",
		"november",
		"desember",
	}
	
	var slice1 = bulan[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice1[0] = "Mei lagi"
	fmt.Println(bulan)
}