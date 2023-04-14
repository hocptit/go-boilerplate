package basedto

type PaginationParams struct {
	Limit  int `json:"-"`
	Offset int `json:"-"`
}

type PaginationDto struct {
	PageSize int `json:"pageSize" form:"pageSize" binding:"max=1000,min=1"`
	Page     int `json:"page" form:"page" binding:"max=10000,min=1"`
	// todo: if page gt 10000 => should implement cursor and pointer
	Cursor     string `json:"cursor" form:"cursor"`
	NextCursor string `json:"next_cursor" form:"next_cursor"`
	PaginationParams
}
