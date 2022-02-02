package models 

type Book struct {
	ID uint `json:"id"  gorm:"primary_key"`
	Title string `json:"title"`
	Author string `json:"author"`
}

//validates input data when creating a new book record
type CreateBookInput struct {
	Title string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

// validates input data when updating a new book record
type UpdateBookInput struct {
	ID uint 		`json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
}