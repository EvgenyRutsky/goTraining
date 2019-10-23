package customarrays

import (
	"fmt"
)

//CreatArray ...
func creatArray(size int, symbol string) [][]string {

	array := make([][]string, size)

	for i := 0; i < size; i++ {
		array[i] = make([]string, size)
		for j := 0; j < size; j++ {
			array[i][j] = symbol
		}
	}

	return array

}

//FillDiagonals ...
func fillDiagonals(slice [][]string, symbol string) [][]string {

	array := slice
	j := 0
	size := len(array)
	for i := 0; i < size; i++ {
		array[i][j] = symbol
		array[size-j-1][i] = symbol
		j++
	}
	return array
}

//PrintSlice ...
func PrintSlice(size int, symbol1 string, symbol2 string) {

	slice := fillDiagonals(creatArray(size, symbol1), symbol2)

	for _, val := range slice {
		var s string
		for _, val1 := range val {
			s = s + val1
		}
		fmt.Println(s)
	}

}
