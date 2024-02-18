BEGIN;

CREATE TABLE decks
(
  id         INT AUTO_INCREMENT PRIMARY KEY,
  game_id    INT      NOT NULL,
  card_id    INT      NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  deleted_at DATETIME
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

COMMIT;
