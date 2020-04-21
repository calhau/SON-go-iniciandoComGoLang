package main

import "fmt"

func main() {
	x := 10
	fmt.Println(&x) //0xc0000120b0

	y := &x
	fmt.Println(y)  //0xc0000120b0
	fmt.Println(*y) //10

	*y = 20
	fmt.Println(x) //20

	var z *int = &x
	fmt.Println(z)  //0xc0000120b0
	fmt.Println(*z) //20

	b := 10
	xpto(b)
	fmt.Println(xpto(b)) //11
	fmt.Println((b))     //10

	fmt.Println(xpto_novo(&b)) //100
	fmt.Println(b)             //100

}

func xpto(a int) int {
	a = a + 1
	return a
}

func xpto_novo(a *int) int {
	*a = 100
	return *a
}
