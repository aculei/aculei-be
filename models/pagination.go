package models

import (
	"strconv"
)

type SortBy int

const (
	SortByDate SortBy = iota
	SortByCam
	SortByAnimal
	SortByTemperature
	SortByMoonPhase
)

var sortByMap = map[SortBy]string{
	SortByDate:        "date",
	SortByCam:         "cam",
	SortByAnimal:      "animal",
	SortByTemperature: "temperature",
	SortByMoonPhase:   "moon_phase",
}

var sortByReverseMap = map[string]SortBy{
	"date":        SortByDate,
	"cam":         SortByCam,
	"animal":      SortByAnimal,
	"temperature": SortByTemperature,
	"moon_phase":  SortByMoonPhase,
}

func (s SortBy) String() string {
	return sortByMap[s]
}

type PaginatedResponseModel[T any] struct {
	Data   []T    `json:"data"`
	SortBy string `json:"sortby"`
	Page   int    `json:"page"`
	Size   int    `json:"size"`
	Count  int    `json:"count"`
	Total  int    `json:"total"`
	Next   *int   `json:"next"`
}

type Paginator struct {
	Size   int    `json:"size"`
	Page   int    `json:"page"`
	SortBy SortBy `json:"sortby"`
}

func NewPaginator(page string, size string, total int, sortBy string) *Paginator {
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}

	countInt, err := strconv.Atoi(size)
	if err != nil {
		countInt = 99999
	}

	sortByEnum, ok := sortByReverseMap[sortBy]
	if !ok {
		sortByEnum = SortByDate
	}

	return &Paginator{
		Size:   countInt,
		Page:   pageInt,
		SortBy: sortByEnum,
	}
}
