package admin

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path"
)

func (a *admin) ResetDB(ctx context.Context) error {
	if err := a.DropDB(ctx); err != nil {
		return err
	}
	if err := a.MigrateDB(ctx); err != nil {
		return err
	}

	if err := a.Seed(ctx); err != nil {
		return err
	}
	fmt.Println("Seed applied!!")

	return nil
}

func (a *admin) DropDB(ctx context.Context) error {
	ddlFile, ok := os.LookupEnv("DROP_DDL_SQL_FILE")
	if !ok {
		return errors.New("DROP_DDL_SQL_FILE is not set")
	}
	ddl, err := os.ReadFile(ddlFile)
	if err != nil {
		return err
	}

	if _, err := a.DB().ExecContext(ctx, string(ddl)); err != nil {
		return err
	}
	fmt.Println("Dropped DB!!")
	return nil
}

func (a *admin) MigrateDB(ctx context.Context) error {
	ddlFile, ok := os.LookupEnv("MIGRATE_DDL_SQL_FILE")
	if !ok {
		return errors.New("MIGRATE_DDL_SQL_FILE is not set")
	}
	ddl, err := os.ReadFile(ddlFile)
	if err != nil {
		return err
	}

	if _, err := a.DB().ExecContext(ctx, string(ddl)); err != nil {
		return err
	}
	fmt.Println("Migrated DB!!")
	return nil
}

func (a *admin) Seed(ctx context.Context) error {
	masterDataDirPath, ok := os.LookupEnv("MASTER_SEED_SQL_DIR_PATH")
	if !ok {
		return errors.New("MASTER_SEED_SQL_DIR_PATH is not set")
	}
	if err := a.applySeed(ctx, masterDataDirPath); err != nil {
		return err
	}
	fmt.Println("")
	fmt.Println("Applied master seed!!")
	fmt.Println("")

	localDataDirPath, ok := os.LookupEnv("LOCAL_SEED_SQL_DIR_PATH")
	if !ok {
		return errors.New("LOCAL_SEED_SQL_DIR_PATH is not set")
	}
	if err := a.applySeed(ctx, localDataDirPath); err != nil {
		return err
	}

	fmt.Println("")
	fmt.Println("Applied local seed!!")
	fmt.Println("")

	return nil
}

func (a *admin) applySeed(ctx context.Context, dirPath string) error {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.IsDir() {
			return fmt.Errorf("%s is not a file", file.Name())
		}
		filePath := path.Join(dirPath, file.Name())
		sql, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}

		if _, err := a.DB().ExecContext(ctx, string(sql)); err != nil {
			return fmt.Errorf("failed to apply seed: %s, err: %w", filePath, err)
		}
		fmt.Println("Applied seed: ", filePath)
	}
	return nil
}

// 生のSQLを実行する
func (a *admin) Exec(ctx context.Context, sql string) error {
	res, err := a.DB().QueryContext(ctx, sql)
	if err != nil {
		return err
	}
	fmt.Println("Executed SQL: ", sql)
	fmt.Println("Result: ", *res)
	return nil
}
