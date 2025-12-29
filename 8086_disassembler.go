package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	file, error := os.Open(args[0])
	check(error)
	b := make([]byte, 2)
	_, error = file.Read(b)
	check(error)

	b7 := (b[0] & 0b10000000) >> 7
	b6 := (b[0] & 0b01000000) >> 6
	b5 := (b[0] & 0b00100000) >> 5
	b4 := (b[0] & 0b00010000) >> 4
	b3 := (b[0] & 0b00001000) >> 3
	b2 := (b[0] & 0b00000100) >> 2
	b1 := (b[0] & 0b00000010) >> 1
	b0 := (b[0] & 0b00000001) >> 0

	b15 := (b[1] & 0b10000000) >> 7
	b14 := (b[1] & 0b01000000) >> 6
	b13 := (b[1] & 0b00100000) >> 5
	b12 := (b[1] & 0b00010000) >> 4
	b11 := (b[1] & 0b00001000) >> 3
	b10 := (b[1] & 0b00000100) >> 2
	b09 := (b[1] & 0b00000010) >> 1
	b08 := (b[1] & 0b00000001) >> 0

	fmt.Printf("%b", b7)
	fmt.Printf("%b", b6)
	fmt.Printf("%b", b5)
	fmt.Printf("%b", b4)
	fmt.Printf("%b", b3)
	fmt.Printf("%b", b2)
	fmt.Printf("%b", b1)
	fmt.Printf("%b", b0)

	fmt.Print(" ")

	fmt.Printf("%b", b15)
	fmt.Printf("%b", b14)
	fmt.Printf("%b", b13)
	fmt.Printf("%b", b12)
	fmt.Printf("%b", b11)
	fmt.Printf("%b", b10)
	fmt.Printf("%b", b09)
	fmt.Printf("%b", b08)

	///////////////////////////////////////////
	// for _, v := range b {		 //
	// 	fmt.Printf("%08b ", v)		 //
	// }					 //
	// fmt.Println()			 //
	///////////////////////////////////////////
}

func check(error error) {
	if error != nil {
		panic(error)
	}
}
