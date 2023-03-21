package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	columns := make([][]uint8, 0, dy)
	for y := 0; y < dy; y++ {
		column := make([]uint8, 0, dx)
		for x := 0; x < dx; x++ {
			column = append(column, uint8(Frag(x, y)))
		}
		columns = append(columns, column)
	}
	return columns
}

func Frag(x, y int) uint8 {
	return uint8((x + y) / 2)
}

func main() {
	pic.Show(Pic)
}
