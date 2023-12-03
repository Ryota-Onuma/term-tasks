package tasks

import (
	"context"
	"errors"
	"fmt"

	"github.com/Ryota-Onuma/term-tasks/db/generated/queries"
	"github.com/Ryota-Onuma/term-tasks/src/ui/list"
	"github.com/Ryota-Onuma/term-tasks/src/ui/progress"
	"github.com/Ryota-Onuma/term-tasks/src/ui/table"
	"github.com/Ryota-Onuma/term-tasks/src/ui/textarea"
	"github.com/Ryota-Onuma/term-tasks/src/ui/textinput"
	"github.com/samber/lo"
)

func (t *task) Edit(ctx context.Context) error {
	selectedRow, err := t.list(ctx, "Choose the task you want to edit.")
	if err != nil {
		return err
	}
	if selectedRow.ID == 0 {
		return nil
	}

	title := textinput.New()
	title.SetText(selectedRow.Title)
	if err := title.Open("Edit task title: "); err != nil {
		return err
	}

	detail := textarea.New()
	detail.SetText(selectedRow.Detail)
	if err := detail.Open("Edit task detail: "); err != nil {
		return err
	}

	status := list.New()
	statuses, err := t.Query().ListMasterTaskStates(ctx)
	if err != nil {
		return err
	}
	statusMap := map[string]string{}
	lo.ForEach(statuses, func(s queries.MasterTaskState, _ int) {
		statusMap[s.Label] = s.Value
	})

	itemsForStates := lo.Map(statuses, func(s queries.MasterTaskState, _ int) list.Item { return list.NewItem(s.Label) })

	status.SetText(selectedRow.TaskStateLabel)
	if err := status.Open("Enter task status: ", itemsForStates); err != nil {
		return err
	}
	if _, ok := statusMap[status.Text()]; !ok {
		return errors.New("Invalid status" + status.Text())
	}

	priority := list.New()
	priorities, err := t.Query().ListMasterTaskPriorities(ctx)
	if err != nil {
		return err
	}

	priorityMap := map[string]string{}
	lo.ForEach(priorities, func(p queries.MasterTaskPriority, _ int) {
		priorityMap[p.Label] = p.Value
	})

	itemsForPriorities := lo.Map(priorities, func(p queries.MasterTaskPriority, _ int) list.Item { return list.NewItem(p.Label) })
	priority.SetText(selectedRow.TaskPriorityLabel)
	if err := priority.Open("Enter task priority: ", itemsForPriorities); err != nil {
		return err
	}
	if _, ok := priorityMap[priority.Text()]; !ok {
		return errors.New("Invalid priority" + priority.Text())
	}

	tbl := table.New()
	columns := []table.Column{
		{Title: "Title", Width: 20},
		{Title: "Detail", Width: 40},
		{Title: "Status", Width: 10},
		{Title: "Priority", Width: 20},
	}
	rows := []table.Row{
		{
			Index: 0,
			Body:  []string{title.Text(), detail.Text(), status.Text(), priority.Text()},
		},
	}
	if err := tbl.Open("Changes are blow. Confirm and go ahead.", columns, rows); err != nil {
		return nil
	}
	if !tbl.IsOK() {
		fmt.Println("Canceled")
		return nil
	}

	if _, err := t.Query().UpdateTask(ctx, queries.UpdateTaskParams{
		ID:       selectedRow.ID,
		Title:    title.Text(),
		Detail:   detail.Text(),
		Status:   statusMap[status.Text()],
		Priority: priorityMap[priority.Text()],
	}); err != nil {
		return err
	}

	// 雰囲気を出すためのProgressBar。実際には何もしていない。
	progressBar := progress.New()
	if err := progressBar.Open(); err != nil {
		return errors.New("something went wrong in progress bar, but don't worry. Task is already added to the database")
	}
	return nil
}
