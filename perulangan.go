package main

import "fmt"

func main() {
	//perulangan for
	for i := 0; i < 5; i++ {
		fmt.Println("Angka ", i)
	}

	//perulangan seperti while
	var i = 0
	for i < 5 {
		fmt.Println("Angka", i)
		i++
	}

	//perulangan for true, berhenti dengan break;
	// for {
	// 	fmt.Println("Angka", i)
	// 	i++
	// 	if i == 5 {
	// 		break
	// 	}
	// }

	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}
}
