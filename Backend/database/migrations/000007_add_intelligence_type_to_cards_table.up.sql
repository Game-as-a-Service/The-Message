BEGIN;

ALTER TABLE cards ADD COLUMN intelligence_type TINYINT COMMENT '1 密電, 2 直達, 3 文件' AFTER color;

COMMIT;
