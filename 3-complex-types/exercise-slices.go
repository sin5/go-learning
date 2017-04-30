package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	xy := make([][]uint8, dx)

	for i := 0; i < dx; i++ {
		xy[i] = make([]uint8, dy)
		for j := 0; j < dy; j++ {
			xy[i][j] = uint8((i + j) / 2)
		}
	}
	return xy
}

func main() {
	pic.Show(Pic)
}
