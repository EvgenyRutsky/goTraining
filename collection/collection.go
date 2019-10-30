package collection

type Element struct {
	index    int
	value    string
	previous *Element
	next     *Element
}

type Collection struct {
	name    string
	first   *Element
	last    *Element
	current *Element
}

func createCollection(name string) *Collection {
	return &Collection{
		name: name,
	}
}

func (collection *Collection) Add(element string) {

}
