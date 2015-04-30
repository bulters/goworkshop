package main

import "fmt"

func main() {
	var i int32 = 1
	var n *int32 = &i

	fmt.Println("i: ", i)
	fmt.Println("n: ", n)
	fmt.Println("*n:", *n)
	fmt.Println("Before", i, *n)

	i = 2
	fmt.Println("After 1", i, *n)

	// n = 3 geeft compilatie fout: // HL
	// cannot use 3 (type int) as type *int32 in assignment //HL
	*n = 3
	fmt.Println("After 2", i, *n)
}
