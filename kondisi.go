package main

import "fmt"

func mainq() {
	var point = 8

	if point == 10 {
		fmt.Println("nilai 10")
	} else if point > 5 {
		fmt.Println("nilai lebih dari 5")
	} else {
		fmt.Println("bukan semua nilai")
	}

	//switch
	switch point {
	case 8:
		fmt.Println("perfect")
	case 7, 9:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}
}
