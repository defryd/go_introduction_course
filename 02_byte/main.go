package main

import "fmt"

func main() {
	var A byte = 'A'
	var a byte = 'a'
	var R byte = 82
	var s byte = 115

	// crear un vector
	vector := []byte{65, 97, 82, 115}

	fmt.Println()
	fmt.Println("Value in ASCII Code:")
	fmt.Println("A:", A)
	fmt.Println("B:", a)
	fmt.Println("R:", R)
	fmt.Println("s:", s)
	fmt.Println("vector", vector)

	fmt.Println()
	fmt.Println("Value in String:")
	fmt.Println("A:", string(A))
	fmt.Println("B:", string(a))
	fmt.Println("R:", string(R))
	fmt.Println("s:", string(s))
	fmt.Println("vector", string(vector))

}
