BEGIN;

CREATE TABLE player_cards
(
  id         INT AUTO_INCREMENT PRIMARY KEY,
  player_id  INT          NOT NULL,
  card_id    INT          NOT NULL,
  type       VARCHAR(255) NOT NULL,
  created_at DATETIME     NOT NULL,
  updated_at DATETIME     NOT NULL,
  deleted_at DATETIME,
  FOREIGN KEY (player_id) REFERENCES players (id),
  FOREIGN KEY (card_id) REFERENCES cards (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

COMMIT;
