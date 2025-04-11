package models

import "errors"

type QueryParams struct {
	Limit  int `query:"limit" default:"10"`
	Offset int `query:"offset" default:"0"`
}

func (q *QueryParams) Validate() error {
	if q.Limit < 1 || q.Limit > 100 {
		return errors.New("limit must be between 1 and 100")
	}
	return nil
}

func (q *QueryParams) GetOffset() int {
	return q.Offset * q.Limit
}
