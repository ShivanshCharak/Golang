package main

	

import (
    "fmt"
    "slices"
)

func main() {
	array := [...]int{1, 2, 3, 4, 5}
	for i := 0; i < 10; i++ {
		switch array[i] {
		case 1:
			fmt.Println("output is", 1)
		case 2:
			fmt.Println("output is", 2)
		case 3:
			fmt.Println("output is", 3)
		}
	}
	slice := [...]
}
