-- tasks 
CREATE TABLE IF NOT EXISTS tasks (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  title TEXT NOT NULL,
  detail TEXT NOT NULL DEFAULT '',
  status TEXT NOT NULL,
  priority TEXT NOT NULL,
  created_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%d %H:%M:%f', 'now', 'localtime')),
  updated_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%d %H:%M:%f', 'now', 'localtime')),
  dead_line_at TEXT,
  CONSTRAINT tasks_status_master_task_states_value_fk FOREIGN KEY (status) REFERENCES master_task_states(value),
  CONSTRAINT tasks_priority_master_task_priorities_value_fk FOREIGN KEY (priority) REFERENCES master_task_priorities(value)
);

CREATE TRIGGER IF NOT EXISTS set_tasks_updated_at
BEFORE UPDATE ON tasks
BEGIN
  UPDATE tasks SET updated_at = strftime('%Y-%m-%d %H:%M:%f', 'now', 'localtime') WHERE id = OLD.id;
END;

-- master_task_states
CREATE TABLE IF NOT EXISTS master_task_states (
  value TEXT PRIMARY KEY,
  label TEXT NOT NULL,
  display_order INTEGER NOT NULL
);

-- task_tags
CREATE TABLE IF NOT EXISTS task_tags (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  task_id INTEGER NOT NULL,
  tag_id INTEGER NOT NULL,
  created_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%d %H:%M:%f', 'now', 'localtime')),
  updated_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%d %H:%M:%f', 'now', 'localtime')),
  CONSTRAINT task_tags_task_id_tasks_id_fk FOREIGN KEY (task_id) REFERENCES tasks(id),
  CONSTRAINT task_tags_tag_id_tags_id_fk FOREIGN KEY (tag_id) REFERENCES tags(id)
);

CREATE TRIGGER IF NOT EXISTS set_tasks_tags_updated_at
BEFORE UPDATE ON task_tags
BEGIN
  UPDATE task_tags SET updated_at = strftime('%Y-%m-%d %H:%M:%f', 'now', 'localtime') WHERE id = OLD.id;
END;

-- tags
CREATE TABLE IF NOT EXISTS tags (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  created_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%d %H:%M:%f', 'now', 'localtime')),
  updated_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%d %H:%M:%f', 'now', 'localtime'))
);

CREATE TRIGGER IF NOT EXISTS set_tags_updated_at
BEFORE UPDATE ON tags
BEGIN
  UPDATE tags SET updated_at = strftime('%Y-%m-%d %H:%M:%f', 'now', 'localtime') WHERE id = OLD.id;
END;

-- master_task_priorities
CREATE TABLE IF NOT EXISTS master_task_priorities (
  value TEXT PRIMARY KEY,
  label TEXT NOT NULL,
  display_order INTEGER NOT NULL
);
