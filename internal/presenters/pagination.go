package presenters

type PresenterPagination[T any] struct {
	Page    *int  `json:"page"`
	PerPage *int  `json:"per_page"`
	Total   int64 `json:"total"`
	Items   []T   `json:"items"`
}

type PresenterPaginationInQuery struct {
	Page    *int `form:"page"`
	PerPage *int `form:"per_page"`
}
