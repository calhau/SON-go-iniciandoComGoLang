package main

import "fmt"

func main() {
	a := "Bruno1"

	switch a {
	case "Ana":
		fmt.Println("Nome é Ana")
	case "Rodrigo":
		fmt.Println("Nome é Rodrigo")
	case "Bruno":
		fmt.Println("Nome é Bruno")
	default:
		fmt.Println("Nome desconhecido!")
	}
}
