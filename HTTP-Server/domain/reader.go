package domain

type Reader struct {
	ID int `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
	BookID int `json:"book_id" bson:"book_id"`
}
