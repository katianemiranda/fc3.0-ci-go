package main

import "fmt"

//import "fmt"

func main() {
	fmt.Println(soma(20, 10))
	fmt.Println(soma(40, 10))

}

//Soma - realiza a soma de dois números
func soma(a int, b int) int {
	return a + b
}
