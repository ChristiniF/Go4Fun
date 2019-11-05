package main

import "fmt"

func sumTheSlice(s []int) (smin int, smax int) {
	min := s[0]
	max := s[0]
	var sum int
	for _, v := range s {

		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
		sum = sum + v
	}
	fmt.Println(min, max, sum)
	smin = sum - min
	smax = sum - max
	return smin, smax
}

func main() {
	var pow = []int{-3, 2, 4, 8, 16000, 32, 64, 128}
	sminn, smaxx := sumTheSlice(pow)
	fmt.Println(sminn, smaxx)

}