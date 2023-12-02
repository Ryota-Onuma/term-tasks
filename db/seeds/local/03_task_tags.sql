INSERT INTO task_tags (task_id, tag_id) 
VALUES
  (1, 1),
  (1, 2),
  (2, 1),
  (2, 3),
  (3, 1),
  (3, 2),
  (3, 3)
ON CONFLICT DO UPDATE SET
  task_id = EXCLUDED.task_id, tag_id = EXCLUDED.tag_id;
