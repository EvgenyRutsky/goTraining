package list

//Element structure
type Element struct {
	value    string
	previous *Element
	next     *Element
}

//Next function returns the next Element
func (element *Element) Next() *Element {

	if element.next == nil {
		return nil
	}

	return element.next

}

//Prev function returns the previous Element
func (element *Element) Prev() *Element {
	if element.previous == nil {
		return nil
	}

	return element.previous

}

//Value function returns value stored in the Element
func (element *Element) Value() string {
	return element.value
}
