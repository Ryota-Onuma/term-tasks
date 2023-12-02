-- name: GetTask :one
SELECT tasks.*, tags.name AS tag_name FROM tasks
JOIN master_task_states ON tasks.status = master_task_states.value
JOIN master_task_priorities ON tasks.priority = master_task_priorities.value
LEFT OUTER JOIN task_tags ON tasks.id = task_tags.task_id
JOIN tags ON task_tags.tag_id = tags.id
WHERE tasks.id = ? LIMIT 1;

-- name: ListTasks :many
SELECT 
  tasks.*, 
  master_task_states.label AS task_state_label, 
  master_task_priorities.label AS task_priority_label,
  tags.name AS tag_name 
FROM tasks
INNER JOIN master_task_states ON tasks.status = master_task_states.value
INNER JOIN master_task_priorities ON tasks.priority = master_task_priorities.value
LEFT OUTER JOIN task_tags ON tasks.id = task_tags.task_id
LEFT OUTER JOIN tags ON task_tags.tag_id = tags.id
ORDER BY master_task_priorities.display_order ASC, tasks.created_at DESC;
;

-- name: CreateTask :one
INSERT INTO tasks (
  title,
  detail,
  status,
  priority,
  dead_line_at
) VALUES (
  ?,
  ?,
  ?,
  ?,
  ?
) RETURNING *;

-- name: UpdateTask :one
UPDATE tasks
SET title = ?,
    detail = ?,
    status = ?,
    priority = ?,
    dead_line_at = ?
WHERE id = ?
RETURNING *;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = ?;
