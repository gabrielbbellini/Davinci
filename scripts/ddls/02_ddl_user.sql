CREATE TABLE IF NOT EXISTS user
(
    id                      INTEGER       NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name                    VARCHAR(60)   NOT NULL DEFAULT '',
    email                   VARCHAR(60)   NOT NULL,
    password                TEXT          NOT NULL,
    role                    TEXT          NOT NULL,
    status_code             TINYINT       NOT NULL DEFAULT 0,
    created_at              DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at             DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);