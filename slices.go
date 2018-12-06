package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	ya := make([][]uint8, dy)
	for y := range ya {
		xa := make([]uint8, dx)
		for x := range xa {
			// fmt.Println(x, y)
			xa[x] = uint8((x + y) / 2)
		}
		ya[y] = xa
	}

	return ya
}

func main() {
	pic.Show(Pic)
}
