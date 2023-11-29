BEGIN;

CREATE TABLE players
(
    id            INT AUTO_INCREMENT PRIMARY KEY,
    game_id       INT          NOT NULL,
    name          VARCHAR(255) NOT NULL,
    identity_card VARCHAR(255) NOT NULL,
    created_at    DATETIME     NOT NULL,
    updated_at    DATETIME     NOT NULL,
    deleted_at    DATETIME,
    FOREIGN KEY (game_id) REFERENCES games (id)
) ENGINE=InnoDB;

COMMIT;
