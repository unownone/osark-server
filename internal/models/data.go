package models

// GetPaginatedResponse is a response for a paginated request
type GetPaginatedResponse[T any] struct {
	Total  int `json:"total"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Data   []T `json:"data"`
}
