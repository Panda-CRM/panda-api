package helpers

import (
	"math"
)

type Pagination struct {
	ItemPerPage int 	`json:"per_page"`
	TotalPages 	int		`json:"total_pages"`
	StartIndex	int 	`json:"-"`
	Page 		int		`json:"-"`
}

const ITEM_PER_PAGE = 50

func MakePagination(totalItems int, page int, itemPerPage int) Pagination {

	if itemPerPage <= 0 {
		itemPerPage = ITEM_PER_PAGE
	}

	if page <= 0 {
		page = 1
	}

	result := int(math.Ceil(float64(totalItems) / float64(itemPerPage)))

	var pag Pagination
	pag.ItemPerPage = itemPerPage
	pag.Page = page
	pag.TotalPages = result
	pag.StartIndex = (page * itemPerPage) - itemPerPage

	return pag
}