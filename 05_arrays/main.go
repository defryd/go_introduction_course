package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("ARRAYS")
	var myIntVar int
	fmt.Println("myIntVar:", myIntVar)
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myIntVar, unsafe.Sizeof(myIntVar), unsafe.Sizeof(myIntVar)*8)

	var myArrayVar [6]int
	fmt.Println("myArrayVar:", myArrayVar)

	myArrayVar1 := [3]string{"value1", "value2", "value3"}
	fmt.Println("myArrayVar1:", myArrayVar1)

	myArrayVar[1] = 2
	myArrayVar[2] = 5
	myArrayVar[3] = 9
	fmt.Println("myArrayVar:", myArrayVar)

	fmt.Println("Size of myArrayVar:", len(myArrayVar))
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myArrayVar, unsafe.Sizeof(myArrayVar), unsafe.Sizeof(myArrayVar)*8)

	fmt.Println()
	fmt.Println("SLICE")

	// Un slice es un puntero a un array
	var slice1 []int
	fmt.Printf("size of slice1: %d, value %v\n", len(slice1), slice1)

	slice1 = append(slice1, 10, 20, 30, 40, 50)
	fmt.Printf("size of slice1: %d, value %v\n", len(slice1), slice1)

	//toma los valores del 5 al 9 de myArrayVar
	mySlice := myArrayVar[2:4]
	fmt.Println(mySlice)
	fmt.Println("Size of mySlice: ", len(mySlice))

	fmt.Println(&myArrayVar[2])
	fmt.Println(&mySlice[0])

	fmt.Println("myArrayVar antes ", myArrayVar)
	fmt.Println("mySlice antes ", mySlice)

	mySlice[0] = 80
	mySlice[1] = 100

	fmt.Println("myArrayVar despues ", myArrayVar)
	fmt.Println("mySlice despues ", mySlice)

	fmt.Println("del inicio al indc 4", myArrayVar[:4])
	fmt.Println("del indc 2 al final", myArrayVar[2:])

	// otra forma de instanciar un slice
	slice := make([]int, 3)
	fmt.Println(slice)
	fmt.Printf("size: %d, value %v\n", len(slice), slice)
}
