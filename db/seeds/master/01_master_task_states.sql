INSERT INTO master_task_states (value, label, display_order)
VALUES ('TODO', '未着手', 1),
       ('IN_PROGRESS', '進行中', 2),
       ('DONE', '完了', 3)
ON CONFLICT (value) DO UPDATE
SET value = excluded.value,
    label = excluded.label;
