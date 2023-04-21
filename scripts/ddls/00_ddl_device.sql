CREATE TABLE IF NOT EXISTS device
(
    id                      INTEGER       NOT NULL AUTO_INCREMENT PRIMARY KEY,
    id_user                 INTEGER       NOT NULL REFERENCES davinci.user (id),
    name                    VARCHAR(60)   NOT NULL DEFAULT '',
    id_resolution           INTEGER       NOT NULL,
    id_orientation          INTEGER       NOT NULL DEFAULT 0,
    status_code             TINYINT       NOT NULL DEFAULT 0,
    created_at              DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at             DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

ALTER TABLE device
    ADD CONSTRAINT user_fk
        FOREIGN KEY (id_user) REFERENCES davinci.user (id);
