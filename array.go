package main

import "fmt"

func main() {
	//vertikal array
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	fmt.Println("jumlah element ", len(fruits))
	fmt.Println("jumlah element ", fruits)

	//tanpa panjang array
	var numbers = [...]int{2, 4, 5}
	fmt.Println("data array \t:", numbers)

	//multi dimensi
	var numbers1 = [2][3]int{{3, 2, 4}, {4, 5, 3}}
	fmt.Println("multy dimensi", numbers1)

	//array for
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("elemen %d : %s\n", i, fruits[i])
	}

	//array for range
	for i, fruit := range fruits {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}

	//penggunaan variabel underscore _
	for _, fruit := range fruits {
		fmt.Printf("nama buah : %s\n", fruit)
	}
}
