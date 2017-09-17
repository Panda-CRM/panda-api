package helpers

import (
	"math"
	"strconv"
)

type PageParams struct {
	Page        int `json:"page" url:"page,omitempty"`
	ItemPerPage int `json:"per_page" url:"per_page,omitempty"`
	TotalPages  int `json:"total_pages" url:"-"`
	StartIndex  int `json:"-" url:"-"`
}

const ITEM_PER_PAGE = 50

func MakePagination(totalItems int, currentPageI interface{}, itemPerPageI interface{}) PageParams {
	var currentPage, itemPerPage int
	var err1, err2 error
	switch currentPageI.(type) {
	case string:
		currentPage, err1 = strconv.Atoi(currentPageI.(string))
		if err1 != nil {
			currentPage = 1
		}
	case int:
		currentPage = currentPageI.(int)
	}
	switch itemPerPageI.(type) {
	case string:
		itemPerPage, err2 = strconv.Atoi(itemPerPageI.(string))
		if err2 != nil {
			itemPerPage = ITEM_PER_PAGE
		}
	case int:
		itemPerPage = itemPerPageI.(int)
	}

	if itemPerPage <= 0 {
		itemPerPage = ITEM_PER_PAGE
	}
	if currentPage <= 0 {
		currentPage = 1
	}

	return PageParams{
		ItemPerPage: itemPerPage,
		Page:        currentPage,
		TotalPages:  int(math.Ceil(float64(totalItems) / float64(itemPerPage))),
		StartIndex:  (currentPage * itemPerPage) - itemPerPage,
	}
}
