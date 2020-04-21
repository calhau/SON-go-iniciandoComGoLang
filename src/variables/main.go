package main

import "fmt"

func main() {

	a := 10
	b := "Hello"
	c := 10.45
	d := false
	e := 'W'
	//Backtipes serve para colocar varias linhas
	f := `Uouuuu!
	Esse é um exemplo de texto longo!
	`
	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)
	fmt.Printf("%T \n", f)
}
