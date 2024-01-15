BEGIN;

CREATE TABLE games
(
  id                INT AUTO_INCREMENT PRIMARY KEY,
  token             LONGTEXT    NOT NULL,
  status            VARCHAR(10) NOT NULL,
  current_player_id INT,
  created_at        DATETIME    NOT NULL,
  updated_at        DATETIME    NOT NULL,
  deleted_at        DATETIME
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

COMMIT;
