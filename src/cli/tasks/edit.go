package tasks

import (
	"context"

	card "github.com/Ryota-Onuma/todo-app/src/ui/creditcard"
)

func (t *task) Edit(ctx context.Context) error {
	c := card.New()
	if err := c.Open(); err != nil {
		return err
	}
	return nil
}
