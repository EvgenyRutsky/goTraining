package main

import (
	"practice-1/slices2d"
)

func main() {

	arr := slices2d.CreateSlice(5, "0")
	filledArray := slices2d.FillDiagonals(arr, "1")
	slices2d.PrintSlice(filledArray)

}
