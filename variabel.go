package main

import "fmt"

/*
ini komentarku
dan komentarnya
*/
func mainx() {
	//variabel
	var firstName string = "abel"
	//tanpa assigment value
	var lastName string
	lastName = "Amel"
	//tanpa type data
	variabel := "wick"
	//multi variabel
	var one, two string = "satu", "dua"
	three, four, five := "tiga", 4, 2.2
	_ = "sampah"

	//
	var negativeNumber = -1243423644

	fmt.Println("halo", firstName, lastName, variabel)
	fmt.Println(one, two, three, four, five)
	//%d utk format data numerik non-desimal

	//variabel constanta
	const name string = "saep"
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)
	fmt.Println(name)
}
