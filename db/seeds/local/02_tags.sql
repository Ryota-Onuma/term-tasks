INSERT INTO tags (name)
VALUES
  ("開発"), 
  ("仕事"),
  ("日常"),
  ("その他")
ON CONFLICT (id) DO UPDATE
SET name = excluded.name;
