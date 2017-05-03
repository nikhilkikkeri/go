package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	my2d := make([][]uint8, dy)
	for i:=0;i<dy;i++ {
		my2d[i] = make([]uint8,dx)
	}
	for x:=0;x<dx;x++ {
		for y:=0;y<dy;y++ {
			my2d[x][y] = uint8(x*y)
		}
	}
	
	return my2d
			
}

func main() {
	pic.Show(Pic)
}