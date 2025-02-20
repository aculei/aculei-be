package models

import (
	"fmt"
	"strconv"
)

type PaginatedResponseModel[T any] struct {
	Page  int  `json:"page"`
	Size  int  `json:"size"`
	Next  *int `json:"next"`
	Data  []T  `json:"data"`
	Total int  `json:"total"`
	Count int  `json:"count"`
}

type Paginator struct {
	Size int `json:"size"`
	Page int `json:"page"`
}

func NewPaginator(page string, size string, total int) *Paginator {
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		fmt.Printf("Error converting page to int: %v\n", err)
		pageInt = 1
	}

	countInt, err := strconv.Atoi(size)
	if err != nil {
		countInt = 10
	}

	return &Paginator{
		Size: countInt,
		Page: pageInt,
	}
}
