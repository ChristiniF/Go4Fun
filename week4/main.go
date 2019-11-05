package main

import (
	"Go4fun/week4/first"
	"Go4fun/week4/second"
	"Go4fun/week4/third"
	"fmt"
)

func main() {
	var pow1 = []int{-3, 2, -30, 59, 63, -65}
	var pow2 = []int{-3, 2, -30, 59, 63}
	missing1 := first.FindMissingElement(pow1, pow2)
	missing2 := second.FindMissingElement(pow1, pow2)
	missing3 := third.FindMissingElement(pow1, pow2)
	fmt.Println(missing1, missing2, missing3)

}
