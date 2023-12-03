package db

import (
	"context"
	_ "embed"
	"fmt"
)

func (a *admin) ResetDB(ctx context.Context) error {
	if err := a.DropDB(ctx); err != nil {
		return err
	}
	if err := a.MigrateDB(ctx); err != nil {
		return err
	}

	if err := a.SeedMasterData(ctx); err != nil {
		return err
	}

	return nil
}

func (a *admin) DropDB(ctx context.Context) error {
	ddl, err := a.schemaFiles.ReadFile("db/schema/drop.sql")
	if err != nil {
		return err
	}

	if _, err := a.DB().ExecContext(ctx, string(ddl)); err != nil {
		return err
	}
	fmt.Println("✅ DB dropped.")
	fmt.Println("")
	return nil
}

func (a *admin) MigrateDB(ctx context.Context) error {
	ddl, err := a.schemaFiles.ReadFile("db/schema/schema.sql")
	if err != nil {
		return err
	}

	if _, err := a.DB().ExecContext(ctx, string(ddl)); err != nil {
		return err
	}
	fmt.Println("✅ DB migrated.")
	fmt.Println("")
	return nil
}

func (a *admin) SeedMasterData(ctx context.Context) error {
	masterDataSqls, err := a.getMasterDataSqls()
	if err != nil {
		return err
	}
	for _, sql := range masterDataSqls {
		if err := a.applySeed(ctx, sql); err != nil {
			return fmt.Errorf("failed to apply master data: %w", err)
		}
	}

	fmt.Println("✅ Master data applied.")
	fmt.Println("")

	return nil
}

func (a *admin) SeedLocalData(ctx context.Context) error {
	localDataSqls, err := a.getLocalDataSqls()
	if err != nil {
		return err
	}
	for _, sql := range localDataSqls {
		if err := a.applySeed(ctx, sql); err != nil {
			return fmt.Errorf("failed to apply local data: %w", err)
		}
	}

	fmt.Println("✅ Sample data applied.")
	fmt.Println("")

	return nil
}

func (a *admin) Seed(ctx context.Context) error {
	if err := a.SeedMasterData(ctx); err != nil {
		return err
	}

	if err := a.SeedLocalData(ctx); err != nil {
		return err
	}

	fmt.Println("🌱 Seed planted.")
	fmt.Println("")

	return nil
}

func (a *admin) getMasterDataSqls() ([]string, error) {
	files, err := a.masterDataFiles.ReadDir("db/seeds/master")
	if err != nil {
		return nil, err
	}

	var sqls []string
	for _, file := range files {
		sqlByte, err := a.masterDataFiles.ReadFile("db/seeds/master/" + file.Name())
		if err != nil {
			return nil, err
		}
		sqls = append(sqls, string(sqlByte))
	}
	return sqls, nil
}

func (a *admin) getLocalDataSqls() ([]string, error) {
	files, err := a.localDataFiles.ReadDir("db/seeds/local")
	if err != nil {
		return nil, err
	}
	var sqls []string
	for _, file := range files {
		sqlByte, err := a.localDataFiles.ReadFile("db/seeds/local/" + file.Name())
		if err != nil {
			return nil, err
		}
		sqls = append(sqls, string(sqlByte))
	}
	return sqls, nil
}

func (a *admin) applySeed(ctx context.Context, sql string) error {
	if _, err := a.DB().ExecContext(ctx, sql); err != nil {
		return err
	}
	return nil
}
