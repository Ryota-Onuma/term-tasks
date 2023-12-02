package tasks

import (
	"context"
	"fmt"

	"github.com/Ryota-Onuma/todo-app/src/ui/table"
)

func (t *task) List(ctx context.Context) error {
	tasks, err := t.Query().ListTasks(ctx)
	if err != nil {
		return err
	}

	tbl := table.New()
	columns := []table.Column{
		{Title: "Title", Width: 20},
		{Title: "Detail", Width: 20},
		{Title: "Status", Width: 10},
		{Title: "Priority", Width: 20},
	}

	var rows []table.Row
	for _, task := range tasks {
		rows = append(rows, table.Row{
			task.Title,
			task.Detail,
			task.TaskStateLabel,
			task.TaskPriorityLabel,
		})
	}

	if err := tbl.Open("ALL TASKS", columns, rows); err != nil {
		return nil
	}
	if !tbl.IsOK() {
		fmt.Println("Canceled")
		return nil
	}

	return nil
}
