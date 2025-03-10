package models

import (
	"strconv"
)

type PaginatedResponseModel[T any] struct {
	Data  []T  `json:"data"`
	Page  int  `json:"page"`
	Size  int  `json:"size"`
	Count int  `json:"count"`
	Total int  `json:"total"`
	Next  *int `json:"next"`
}

type Paginator struct {
	Size int `json:"size"`
	Page int `json:"page"`
}

func NewPaginator(page string, size string, total int) *Paginator {
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}

	countInt, err := strconv.Atoi(size)
	if err != nil {
		countInt = 99999
	}

	return &Paginator{
		Size: countInt,
		Page: pageInt,
	}
}
