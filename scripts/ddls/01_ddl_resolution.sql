CREATE TABLE IF NOT EXISTS resolution
(
    id                      INTEGER       NOT NULL AUTO_INCREMENT PRIMARY KEY,
    width                   INTEGER       NOT NULL,
    height                  INTEGER       NOT NULL,
    status_code             TINYINT       NOT NULL DEFAULT 0,
    created_at              DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at             DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);