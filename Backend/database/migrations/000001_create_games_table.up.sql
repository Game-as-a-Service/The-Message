BEGIN;

CREATE TABLE games
(
    id         INT AUTO_INCREMENT PRIMARY KEY,
    token      LONGTEXT NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME
) ENGINE=InnoDB;

COMMIT;
