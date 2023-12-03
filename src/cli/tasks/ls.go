package tasks

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Ryota-Onuma/terminal-task-manager/db/generated/queries"
	"github.com/Ryota-Onuma/terminal-task-manager/src/ui/table"
)

func (t *task) List(ctx context.Context) error {
	_, err := t.list(ctx, "Search results are below.")
	if err != nil {
		return err
	}
	return nil
}

func (t *task) list(ctx context.Context, tableTitle string) (queries.ListTasksRow, error) {
	tasks, err := t.Query().ListTasks(ctx)
	if err != nil {
		return queries.ListTasksRow{}, err
	}

	tbl := table.New()
	columns := []table.Column{
		{Title: "ID", Width: 10},
		{Title: "Title", Width: 20},
		{Title: "Detail", Width: 40},
		{Title: "Status", Width: 10},
		{Title: "Priority", Width: 20},
	}

	var rows []table.Row
	for id, task := range tasks {
		rows = append(rows, table.Row{
			Index: id,
			Body: []string{
				strconv.Itoa(int(task.ID)),
				task.Title,
				task.Detail,
				task.TaskStateLabel,
				task.TaskPriorityLabel,
			},
		})
	}

	if err := tbl.Open(tableTitle, columns, rows); err != nil {
		return queries.ListTasksRow{}, err
	}
	if !tbl.IsOK() {
		fmt.Println("Canceled")
		return queries.ListTasksRow{}, nil
	}
	selectedRowIndex := tbl.SelectedRow().Index
	selectedRow := tasks[selectedRowIndex]
	return selectedRow, nil
}
