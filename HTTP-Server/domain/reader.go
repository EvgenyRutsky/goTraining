package domain

type Reader struct {
	ID int `json:"id"`
	Name string `json:"name"`
	BookID int `json:"book_id"`
}
