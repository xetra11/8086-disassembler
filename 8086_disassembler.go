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
	for j := 0; j < 2; j++ {
		for i := 0; i < 8; i++ {
			fmt.Printf("%d", (b[j]>>i)&1)
			if (i == 7) {
				fmt.Print(" ")
			}
		}
	}

}

func check(error error) {
	if error != nil {
		panic(error)
	}
}
