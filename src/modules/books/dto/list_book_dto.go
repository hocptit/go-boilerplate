package dto

type ListBookDto struct {
	Author int    `form:"author" binding:"required"`
	Title  string `form:"title"`
}
