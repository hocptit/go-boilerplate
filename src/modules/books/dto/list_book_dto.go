package dto

type ListBookDto struct {
	Author int    `form:"author"`
	Title  string `form:"title"`
}
