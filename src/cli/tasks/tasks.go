package tasks

import (
	"github.com/Ryota-Onuma/terminal-task-manager/db/generated/queries"
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
