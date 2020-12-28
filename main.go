package main

import "fmt"

type (
	height int
	width  int
	area   int
)

var (
	vWidth, vHeight int
	vArea           area
)

func main() {
	var area = func(pWidth width, pHeight height) area {
		vWidth = int(pWidth)
		vHeight = int(pHeight)
		vArea = area(vHeight * vWidth)
		return vArea
	}

	var inc = func(pWidth *width, pHeight *height) (width, height) {
		*pWidth += 15
		*pHeight += 15
		return *pWidth, *pHeight
	}

	var ww width = 10
	var hh height = 5

	fmt.Println(inc(&ww, &hh))
	fmt.Println(side(inc))
	fmt.Println("With area and increment:", area(inc(side(inc))))

}

func side(fn func(pWidth *width, pHeight *height) (width, height)) (*width, *height) {
	var vHeight height
	var vWidth width
	fmt.Println("Enter width")
	fmt.Scan(&vWidth)
	fmt.Println("Enter height")
	fmt.Scan(&vHeight)

	return &vWidth, &vHeight
}
