INSERT INTO master_task_priorities (value, label, display_order)
VALUES ('URGENT_AND_IMPORTANT', '緊急かつ重要', 1),
       ('URGENT_AND_NOT_IMPORTANT', '緊急だが重要ではない', 2),
       ('NOT_URGENT_AND_IMPORTANT', '緊急ではないが重要', 3),
       ('NOT_URGENT_AND_NOT_IMPORTANT', '緊急でも重要でもない', 4)
ON CONFLICT (value) DO UPDATE
SET value = excluded.value,
    label = excluded.label;
