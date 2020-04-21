package main

import "fmt"

//funcao q recebe 2 parametros e retorna um inteiro
// func funcName(a int, b string) int {}

func testando(a int) int {
	return a * a
}

func namedReturn(a string) (x string) {
	x = a
	return
}

func moreReturns(a string, b string) (string, string) {
	return a, b
}

//Nessa função porro receber diversos valores para X
func variadicFunc(x ...int) int {
	var res int
	for _, v := range x {
		res += v
	}
	return res
}

func main() {
	fmt.Println(testando(10))

	//Atribuindo a x o retorno da funcao "testando" com parametro 5 // Resultado:25
	x := testando(5)
	fmt.Println(x)

	fmt.Println("Imprimindo namedReturn()")
	fmt.Println(namedReturn("Oie sou a namedReturn()"))

	fmt.Println("Imprimindo moreReturn()")
	c, d := moreReturns("Bruno", "Leal")
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(c, d)

	fmt.Println("Imprimindo variadicFun()")
	fmt.Println(variadicFunc(1, 2, 5, 10, 13))
}
