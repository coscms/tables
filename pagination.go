package tables

import (
	"github.com/webx-top/echo"
	"github.com/webx-top/pagination"
)

func NewPagination(ctx echo.Context) *Pagination {
	return &Pagination{
		p: pagination.New(ctx),
	}
}

type Pagination struct {
	p *pagination.Pagination
}
