CREATE TABLE IF NOT EXISTS user
(
    id          INTEGER      NOT NULL AUTO_INCREMENT PRIMARY KEY,
    id_role     INTEGER      NOT NULL DEFAULT 1,
    name        VARCHAR(20)  NOT NULL DEFAULT '',
    email       VARCHAR(50)  NOT NULL DEFAULT '',
    password    VARCHAR(100) NOT NULL DEFAULT '',
    status_code TINYINT      NOT NULL DEFAULT 0,
    created_at  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS role
(
    id          INTEGER     NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name        VARCHAR(20) NOT NULL DEFAULT '',
    status_code TINYINT     NOT NULL DEFAULT 0,
    created_at  DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE user
    ADD CONSTRAINT role_fk
        FOREIGN KEY (id_role) REFERENCES role (id);
