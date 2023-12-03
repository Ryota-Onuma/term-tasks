package tasks

import (
	"context"
	"errors"
	"fmt"

	"github.com/samber/lo"

	"github.com/Ryota-Onuma/todo-app/db/generated/queries"
	"github.com/Ryota-Onuma/todo-app/src/ui/list"
	"github.com/Ryota-Onuma/todo-app/src/ui/progress"
	"github.com/Ryota-Onuma/todo-app/src/ui/table"
	"github.com/Ryota-Onuma/todo-app/src/ui/textarea"
	"github.com/Ryota-Onuma/todo-app/src/ui/textinput"
)

func (t *task) Add(ctx context.Context) error {
	title := textinput.New()
	if err := title.Open("Enter task title: "); err != nil {
		return err
	}
	detail := textarea.New()
	if err := detail.Open("Enter task content below..."); err != nil {
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
		{title.Text(), detail.Text(), status.Text(), priority.Text()},
	}
	if err := tbl.Open("ALL TASKS", columns, rows); err != nil {
		return nil
	}
	if !tbl.IsOK() {
		fmt.Println("Canceled")
		return nil
	}

	if _, err := t.Query().CreateTask(ctx, queries.CreateTaskParams{
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
		return errors.New("Something went wrong in progress bar, but don't worry. Task is already added to the database.")
	}

	return nil
}
