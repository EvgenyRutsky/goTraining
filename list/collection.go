package list

import (
	"fmt"
)

//Collection structure
type Collection struct {
	length int
	first  *Element
	last   *Element
}

// CreateCollection func
func CreateCollection() *Collection {
	return &Collection{}
}

//Add func adds new element to the Collection
func (collection *Collection) Add(element string) {
	elem := &Element{
		value: element,
	}
	if collection.first == nil {
		collection.first = elem
	} else {
		elem.previous = collection.Last()
		collection.Last().next = elem
	}
	collection.last = elem
	collection.length++

}

// Get function returns element by index
func (collection *Collection) Get(index int) *Element {

	elem := collection.First()

	if collection.first == nil {
		return nil
	}

	for i := 0; i < index; i++ {

		if elem.next == nil {
			return nil
		}

		elem = elem.Next()

	}

	return elem

}

// First function returns the first element of the collection
func (collection *Collection) First() *Element {
	if collection.first == nil {
		return nil
	}

	return collection.first

}

//Last function returns the last element of the collection
func (collection *Collection) Last() *Element {
	elem := collection.First()

	if collection.first == nil {
		return nil
	}

	for elem.next != nil {
		elem = elem.Next()
	}

	return elem
}

//Length returns amount of elements in the collection
func (collection *Collection) Length() int {
	return collection.length
}

//Remove function removes element from the collection
func (collection *Collection) Remove(index int) {
	elem := collection.Get(index)

	if elem.previous == nil {
		collection.first = elem.Next()
		elem.Next().previous = nil

	} else if elem.next == nil {
		collection.last = elem.Prev()
		elem.Prev().next = nil
	} else {
		elem.Prev().next = elem.Next()
		elem.Next().previous = elem.Prev()
	}

	collection.length--
}

//Print function prints the values of the collection
func (collection *Collection) Print() {

	for i := 0; i < collection.length; i++ {
		fmt.Println(collection.Get(i).Value())
	}

}
