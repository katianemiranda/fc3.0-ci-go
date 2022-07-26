package main

import "fmt"

//import "fmt"

func main() {
	fmt.Println(Soma(20, 10))
	fmt.Println(Soma(40, 10))

}

//Soma - realiza a soma de dois números
func Soma(a int, b int) int {
	return a + b
}
