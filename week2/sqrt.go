package main

import "fmt"

func sqrt(y float64) float64 {
	x0 := y / 2
	y0 := x0 * x0
	dy0 := 2 * x0
	x := x0 + ((y - y0) / dy0)

	for i := 0; i < 100; i++ {
		x0 = x
		y0 = x0 * x0
		dy0 = 2 * x0
		x = x0 + ((y - y0) / dy0)
		if (y - x*x) > 0.00001 {
			return x
		}

	}
	return x
}

func main() {
	k := sqrt(255)
	fmt.Print(k)

}
