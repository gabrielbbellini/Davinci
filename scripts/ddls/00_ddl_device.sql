CREATE TABLE IF NOT EXISTS device
(
    id                      INTEGER(9)    NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name                    TEXT,
    id_resolution           INTEGER       NOT NULL,
    id_orientation          INTEGER       NOT NULL,
    status_code             TINYINT       NOT NULL DEFAULT 0,
    created_at              DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at             DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
--  FOREIGN KEY (id_resolution) REFERENCES resolution (id),
--  FOREIGN KEY (id_orientation) REFERENCES orientation (id),
);