package cli

import (
	"database/sql"
	"embed"
	"fmt"

	"github.com/Ryota-Onuma/terminal-task-manager/db/generated/queries"
	"github.com/Ryota-Onuma/terminal-task-manager/src/cli/admin"
	"github.com/Ryota-Onuma/terminal-task-manager/src/cli/tasks"

	_ "embed"

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
			registerAdmin(db, schemaFiles, masterDataFiles, localDataFiles),
			registerVersion(),
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

func registerAdmin(db *sql.DB, schemaFiles, masterDataFiles, localDataFiles embed.FS) *c.Command {
	adminTask := admin.New(db, schemaFiles, masterDataFiles, localDataFiles)
	var adminResetDB, adminMigrateDB, adminSeed c.ActionFunc
	adminResetDB = func(cCtx *c.Context) error {
		if err := adminTask.ResetDB(cCtx.Context); err != nil {
			return err
		}
		return nil
	}
	adminMigrateDB = func(cCtx *c.Context) error {
		if err := adminTask.MigrateDB(cCtx.Context); err != nil {
			return err
		}
		return nil
	}
	adminSeed = func(cCtx *c.Context) error {
		if err := adminTask.Seed(cCtx.Context); err != nil {
			return err
		}
		return nil
	}

	return &c.Command{
		Name:  "admin",
		Usage: "manage admin tasks",
		Subcommands: []*c.Command{
			{
				Name:  "db",
				Usage: "manage database",
				Subcommands: []*c.Command{
					{
						Name:   "reset",
						Usage:  "reset database",
						Action: adminResetDB,
					},
					{
						Name:   "migrate",
						Usage:  "migrate database",
						Action: adminMigrateDB,
					},
					{
						Name:   "seed",
						Usage:  "seed database",
						Action: adminSeed,
					},
				},
			},
		},
	}
}
