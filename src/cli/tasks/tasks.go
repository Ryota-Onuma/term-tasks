package tasks

import (
	"github.com/Ryota-Onuma/term-tasks/db/generated/queries"
)

type task struct {
	query *queries.Queries
}

func New(q *queries.Queries) *task {
	return &task{query: q}
}

func (t *task) Query() *queries.Queries {
	return t.query
}
