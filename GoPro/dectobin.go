package main

import "fmt"

func main() {
	fmt.Println("Enter a decimal Number")
	var dec int
	fmt.Scan(&dec)

	fmt.Printf("%d, in binary format = %b\n", dec, dec)
}
