package slices2d

import (
	"fmt"
)

//CreatSlice function creates square matrix and fills it with specified symbol
func CreateSlice(size int, symbol string) [][]string {

	slice := make([][]string, size)

	for i := 0; i < size; i++ {
		slice[i] = make([]string, size)
		for j := 0; j < size; j++ {
			slice[i][j] = symbol
		}
	}

	return slice

}

//FillDiagonals function fills both diagonals of given 2d slice
func FillDiagonals(givenSlice [][]string, symbol string) [][]string {

	slice := givenSlice

	size := len(slice)

	for i := 0; i < size; i++ {
		slice[i][i] = symbol
		slice[size-i-1][i] = symbol
	}
	return slice
}

//PrintSlice function prints given 2d
func PrintSlice(givenSlice [][]string) {

	slice := givenSlice

	for _, value := range slice {
		var s string
		for _, sub_value := range value {
			s = s + sub_value
		}
		fmt.Println(s)
	}

}
