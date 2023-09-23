package model

const PageDefaultLimit = 24

type Page struct {
	Total  int64 `json:"total"`
	Count  int   `json:"count"`
	Offset int   `json:"offset"`
	Limit  int   `json:"limit"`
}

type Response[T any] struct {
	Data *[]T  `json:"data"`
	Meta *Page `json:"meta"`
}

func NewResponse[T any](data *[]T, page *Page) *Response[T] {
	page.Count = len(*data)
	return &Response[T]{
		Data: data,
		Meta: page,
	}
}
