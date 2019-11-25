package main

import (
	"fmt"
	"practice-1/list"
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

	//Practice 2 - Collections
	myCollection := list.CreateCollection()

	myCollection.Add("the first")
	myCollection.Add("the second")
	myCollection.Add("the third")
	myCollection.Add("the fourth")
	myCollection.Add("the fifth")

	myCollection.Print()

	myCollection.Remove(1)

	myCollection.Print()

	fmt.Printf("The value of the element with index 1 is %v\n", myCollection.Get(1).Value())

	fmt.Printf("The length of the collection equals to %v\n", myCollection.Length())

	fmt.Printf("The first element of the collection is %v\n", myCollection.First().Value())

	fmt.Printf("The last element of the collection is %v\n", myCollection.Last().Value())

}
