package main

import "fmt"

func main() {

	fmt.Println("Imprindo valor de i")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("Imprindo valor de X")
	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}

	fmt.Println("Loop com condicional dentro e Break")
	for {
		fmt.Println(x)
		x++
		if x == 100 {
			fmt.Println("Agora vou parar, valor de x:")
			fmt.Println(x)
			break
			//continue pode ser usado tbm
		}
	}

}
