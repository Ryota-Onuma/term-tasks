package cli

import (
	"database/sql"
	"embed"
	"fmt"

	_ "embed"

	"github.com/Ryota-Onuma/term-tasks/db/generated/queries"
	"github.com/Ryota-Onuma/term-tasks/src/cli/db"
	"github.com/Ryota-Onuma/term-tasks/src/cli/tasks"

	c "github.com/urfave/cli/v2"
)

type Cli struct {
	app             *c.App
	schemaFiles     embed.FS
	masterDataFiles embed.FS
	localDataFiles  embed.FS
}

func New(db *sql.DB, schemaFiles, masterDataFiles, localDataFiles embed.FS) *Cli {
	q := queries.New(db)
	return &Cli{
		app:             registerApp(db, q, schemaFiles, masterDataFiles, localDataFiles),
		schemaFiles:     schemaFiles,
		masterDataFiles: masterDataFiles,
		localDataFiles:  localDataFiles,
	}
}

func (c *Cli) Run(args []string) error {
	return c.app.Run(args)
}

func registerApp(db *sql.DB, q *queries.Queries, schemaFiles, masterDataFiles, localDataFiles embed.FS) *c.App {
	return &c.App{
		Commands: []*c.Command{
			registerTasks(q),
			registerDB(db, schemaFiles, masterDataFiles, localDataFiles),
			registerVersion(),
			registerInit(db, schemaFiles, masterDataFiles, localDataFiles),
		},
	}
}

func registerVersion() *c.Command {
	var versionFunc c.ActionFunc = func(cCtx *c.Context) error {
		fmt.Println(version)
		return nil
	}
	return &c.Command{
		Name:   "version",
		Usage:  "show version",
		Action: versionFunc,
	}
}

func registerInit(sqlDB *sql.DB, schemaFiles, masterDataFiles, localDataFiles embed.FS) *c.Command {
	dbAdmin := db.New(sqlDB, schemaFiles, masterDataFiles, localDataFiles)
	return &c.Command{
		Name:  "init",
		Usage: "initialize",
		Action: func(cCtx *c.Context) error {
			if err := dbAdmin.Init(cCtx.Context); err != nil {
				return err
			}
			return nil
		},
	}
}

func registerTasks(q *queries.Queries) *c.Command {
	task := tasks.New(q)
	var taskList, taskAdd, taskEdit c.ActionFunc
	taskList = func(cCtx *c.Context) error {
		if err := task.List(cCtx.Context); err != nil {
			return err
		}
		return nil
	}
	taskAdd = func(cCtx *c.Context) error {
		if err := task.Add(cCtx.Context); err != nil {
			return err
		}
		return nil
	}
	taskEdit = func(cCtx *c.Context) error {
		if err := task.Edit(cCtx.Context); err != nil {
			return err
		}
		return nil
	}

	return &c.Command{
		Name:    "tasks",
		Aliases: []string{"ts"},
		Usage:   "manage tasks",
		Action:  taskList,
		Subcommands: []*c.Command{
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "list tasks",
				Action:  taskList,
			},
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a task",
				Action:  taskAdd,
			},
			{
				Name:    "edit",
				Aliases: []string{"e"},
				Usage:   "edit a task",
				Action:  taskEdit,
			},
		},
	}
}

func registerDB(sqlDB *sql.DB, schemaFiles, masterDataFiles, localDataFiles embed.FS) *c.Command {
	dbAdmin := db.New(sqlDB, schemaFiles, masterDataFiles, localDataFiles)
	var resetDB, migrateDB, seed c.ActionFunc
	resetDB = func(cCtx *c.Context) error {
		if err := dbAdmin.ResetDB(cCtx.Context); err != nil {
			return err
		}
		return nil
	}
	migrateDB = func(cCtx *c.Context) error {
		if err := dbAdmin.MigrateDB(cCtx.Context); err != nil {
			return err
		}
		return nil
	}
	seed = func(cCtx *c.Context) error {
		if err := dbAdmin.Seed(cCtx.Context); err != nil {
			return err
		}
		return nil
	}

	return &c.Command{
		Name:  "db",
		Usage: "manage db",
		Subcommands: []*c.Command{
			{
				Name:   "reset",
				Usage:  "reset database",
				Action: resetDB,
			},
			{
				Name:   "migrate",
				Usage:  "migrate database",
				Action: migrateDB,
			},
			{
				Name:   "seed",
				Usage:  "seed database",
				Action: seed,
			},
		},
	}
}
