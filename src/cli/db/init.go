package db

import (
	"context"
	"fmt"
)

func (a *admin) Init(ctx context.Context) error {
	if err := a.MigrateDB(ctx); err != nil {
		return err
	}

	if err := a.SeedMasterData(ctx); err != nil {
		return err
	}

	fmt.Println("Thanks🙏 Initialization completed.")
	return nil
}
