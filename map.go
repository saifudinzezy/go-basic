package main

import "fmt"

func main() {

	/* var chicken3 = map[string]int{}
	var chicken4 = make(map[string]int)
	var chicken5 = *new(map[string]int) */

	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 1
	chicken["februari"] = 2

	fmt.Println("januari", chicken["januari"])
	//menghapus item map
	delete(chicken, "januari")
	fmt.Println(len(chicken))

	//cara horizontal
	var ayam = map[string]int{"jago": 10, "lehor": 20}
	fmt.Println("ayam", ayam["jago"])

	for key, val := range chicken {
		fmt.Println(key, "  \t:", val)
	}

	//deteksi keberadaan item dengan key tertentu

	var xman = map[string]int{"januari": 50, "februari": 40}
	// var value, isExist = xman["mei"]
	var value, isExist = xman["januari"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}

	var chickens = []map[string]string{
		map[string]string{"name": "chicken blue", "gender": "male"},
		map[string]string{"name": "chicken red", "gender": "male"},
		map[string]string{"name": "chicken yellow", "gender": "female"},
	}

	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}
}
