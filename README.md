# Term Tasks - manage your tasks from your terminal app

## What is Term Tasks ? 
`Term Tasks` is an app that works from your terminal app, made by [@Ryota-Onuma](https://github.com/Ryota-Onuma).
You can create, list, edit, and delete your tasks with easy ways.

※ Some of the features, such as delete tasks, remind and attach tags are unavailable. Those convenients will be introduced to this app in the future.

## Used technoloiges for development
- [Go](https://go.dev/)
- [sqlc](https://sqlc.dev/)
- [SQLite](https://www.sqlite.org/index.html)
- [Bubble Tea](https://github.com/charmbracelet/bubbletea)
- [urfave/cli](https://github.com/urfave/cli)

## 🚀 Install

```sh
$ go install github.com/Ryota-Onuma/term-tasks@latest
```

## 💡 Usage

If you want to know shorthands, try `--help`.

### 🙈 Show version

```sh
$ term-tasks version
v0.0.1
```

### 🙏 Initialization

```sh
$ term-tasks init
```

### 🌱 Apply sample data

If you want to try `Term Tasks`, it's recommended to execute this command.

```sh
$ term-tasks db seed
```

### 😭 DB Reset

※ All your tasks will be deleted.

```sh
$ term-tasks db reset
```

### ✅ Add your task
```sh
$ term-tasks tasks add
```

### 📃 List your tasks
```sh
$ term-tasks tasks list
```

### ✏️ Edit your tasks
```sh
$ term-task tasks edit
```
### 💥 Delete your tasks

Comming soon... 🙏

### 📃 Check a task

Comming soon... 🙏

### 🕐 Set a reminder and deadline

Comming soon... 🙏


