package cli

import (
	"database/sql"

	"github.com/Ryota-Onuma/todo-app/db/generated/queries"
	"github.com/Ryota-Onuma/todo-app/src/cli/admin"
	"github.com/Ryota-Onuma/todo-app/src/cli/tasks"

	c "github.com/urfave/cli/v2"
)

type Cli struct {
	app *c.App
}

func New(db *sql.DB) *Cli {
	q := queries.New(db)
	return &Cli{
		app: registerApp(db, q),
	}
}

func (c *Cli) Run(args []string) error {
	return c.app.Run(args)
}

func registerApp(db *sql.DB, q *queries.Queries) *c.App {
	return &c.App{
		Commands: []*c.Command{
			registerTasks(q),
			registerAdmin(db),
		},
	}
}

func registerTasks(q *queries.Queries) *c.Command {
	task := tasks.New(q)
	var taskList, taskAdd c.ActionFunc
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
		},
	}
}

func registerAdmin(db *sql.DB) *c.Command {
	adminTask := admin.New(db)
	var adminResetDB, adminMigrateDB, adminSeed, execSql c.ActionFunc
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

	execSql = func(cCtx *c.Context) error {
		sql := cCtx.Args().First()
		if err := adminTask.Exec(cCtx.Context, sql); err != nil {
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
					{
						Name:   "exec",
						Usage:  "exec sql",
						Action: execSql,
					},
				},
			},
		},
	}
}
