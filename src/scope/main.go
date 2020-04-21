package main

import "fmt"

var y int = 20

func main() {
	x := 10
	fmt.Println(x)

	//Consigo enxergar pois a variavel Y esta setada no scope do pacote
	fmt.Println(y)

	fmt.Println(z)

	printNome()
}

func printX() {
	//Não consigo ver, pois X ta no scope de main()
	//fmt.Println(x)

	//Esse consigo ver pois esta no scope do pacote
	fmt.Println(y)
}
