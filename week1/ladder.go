package main

import "fmt"

func ladder(n int) {

	if n > 0 {
		spaces := ""
		hash := ""

		for i := 1; i < n+1; i++ {
			for k := 0; k < (n - i); k++ {
				spaces += " "
			}
			for k := 0; k < i; k++ {
				hash += "#"
			}
			fmt.Print(spaces, hash, "\n")
			hash = ""
			spaces = ""
		}

	} else {
		fmt.Println("wrong number")
	}
}

func main() {
	ladder(4)
}
