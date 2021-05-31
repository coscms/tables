package tables

import (
	"github.com/webx-top/echo"
	"github.com/webx-top/pagination"
)

func NewPagination(ctx echo.Context) *Pagination {
	return &Pagination{
		Pagination: pagination.New(ctx),
	}
}

type Pagination struct {
	*pagination.Pagination
}
