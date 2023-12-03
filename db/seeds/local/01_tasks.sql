INSERT INTO tasks (id, title, detail, status, priority, dead_line_at)
VALUES (1, 'タスク1', 'タスク1の詳細', 'TODO', 'URGENT_AND_IMPORTANT','2023-01-01 00:00:00'),
       (2, 'タスク2', 'タスク2の詳細', 'IN_PROGRESS', 'URGENT_AND_IMPORTANT','2024-01-02 00:00:00'),
       (3, 'タスク3', 'タスク3の詳細', 'DONE','URGENT_AND_IMPORTANT','2024-01-03 00:00:00'),
       (4, 'タスク4', 'タスク4の詳細', 'TODO','URGENT_AND_NOT_IMPORTANT','2024-01-04 00:00:00'),
       (5, 'タスク5', 'タスク5の詳細', 'IN_PROGRESS','URGENT_AND_NOT_IMPORTANT','2024-01-05 00:00:00'),
       (6, 'タスク6', 'タスク6の詳細', 'DONE','NOT_URGENT_AND_IMPORTANT','2024-01-06 00:00:00'),
       (7, 'タスク7', 'タスク7の詳細', 'TODO','NOT_URGENT_AND_IMPORTANT','2024-01-07 00:00:00'),
       (8, 'タスク8', 'タスク8の詳細', 'IN_PROGRESS','NOT_URGENT_AND_IMPORTANT','2024-01-08 00:00:00'),
       (9, 'タスク9', 'タスク9の詳細', 'DONE','NOT_URGENT_AND_NOT_IMPORTANT','2024-01-09 00:00:00'),
       (10,'タスク10', 'タスク10の詳細', 'TODO','URGENT_AND_IMPORTANT','2024-01-10 00:00:00'),
       (11,'タスク11', 'タスク11の詳細', 'IN_PROGRESS','URGENT_AND_NOT_IMPORTANT','2024-01-11 00:00:00'),
       (12,'タスク12', 'タスク12の詳細', 'DONE','NOT_URGENT_AND_IMPORTANT','2024-01-12 00:00:00'),
       (13,'タスク13', 'タスク13の詳細', 'TODO','NOT_URGENT_AND_NOT_IMPORTANT','2024-01-13 00:00:00'),
       (14,'タスク14', 'タスク14の詳細', 'IN_PROGRESS','NOT_URGENT_AND_NOT_IMPORTANT','2024-01-14 00:00:00'),
       (15,'タスク15', 'タスク15の詳細', 'DONE','NOT_URGENT_AND_NOT_IMPORTANT','2024-01-15 00:00:00'),
       (16,'タスク16', 'タスク16の詳細', 'TODO','NOT_URGENT_AND_NOT_IMPORTANT','2024-01-16 00:00:00'),
       (17,'タスク17', 'タスク17の詳細', 'IN_PROGRESS','NOT_URGENT_AND_NOT_IMPORTANT','2024-01-17 00:00:00')
ON CONFLICT (id) DO UPDATE
SET title = excluded.title,
    detail = excluded.detail,
    status = excluded.status,
    dead_line_at = excluded.dead_line_at;
