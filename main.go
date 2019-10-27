package main

import (
	"practice-1/randomcube"
	"practice-1/slices2d"
)

func main() {

	//Practice 1, Task 1
	arr := slices2d.CreateSlice(5, "0")
	filledArray := slices2d.FillDiagonals(arr, "1")
	slices2d.PrintSlice(filledArray)

	//Practice 1, Task 2
	randomcube.CountStatistic(1000, 2)

}
