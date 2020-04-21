package main

import "fmt"

func main() {
	a := 2

	if a > 10 {
		fmt.Println("A > 10")
	} else if a > 5 {
		fmt.Println("A > 5")
	} else {
		fmt.Println("A menor que 5")
	}

	if true {
		fmt.Println("Quando for true")
	}
	if false {
		fmt.Println("Quando for false, n imprimi")
	}

	if a >= 10 {
		fmt.Println("A >= 10")
	}

	b := true
	if b {
		fmt.Println("B é true então funciona")
	}

	if x := "Atribui variavel X"; b {
		fmt.Println(x)
	}
}
