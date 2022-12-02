package dto

type CreateBookDto struct {
	Title       string `json:"title" binding:"required"`
	Author      int    `json:"author" binding:"required"`
	Description string `json:"description" binding:"required"`
}
