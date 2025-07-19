package query

import (
	"strconv"
)

const (
	defaultPage     = 1
	defaultPageSize = 10
	maxPageSize     = 100
)

type Pagination struct {
	Page     int
	PageSize int
}

func (p Pagination) Limit() int {
	if p.PageSize <= 0 || p.PageSize > 100 {
		return 10
	}
	return p.PageSize
}

func (p Pagination) Offset() int {
	if p.Page <= 0 {
		return 0
	}
	return (p.Page - 1) * p.Limit()
}

func NewPagination(pageStr, pageSizeStr string) Pagination {
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = defaultPage
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 || pageSize > maxPageSize {
		pageSize = defaultPageSize
	}

	return Pagination{
		Page:     page,
		PageSize: pageSize,
	}
}
