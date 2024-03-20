BEGIN;

CREATE TABLE decks
(
  id         INT AUTO_INCREMENT PRIMARY KEY,
  game_id    INT      NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  deleted_at DATETIME,
  FOREIGN KEY (game_id) REFERENCES games (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE deck_cards
(
  id         INT AUTO_INCREMENT PRIMARY KEY,
  deck_id    INT      NOT NULL,
  card_id    INT      NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  deleted_at DATETIME,
  FOREIGN KEY (deck_id) REFERENCES decks (id),
  FOREIGN KEY (card_id) REFERENCES cards (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

COMMIT;
