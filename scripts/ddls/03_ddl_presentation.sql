CREATE TABLE IF NOT EXISTS presentation
(
    id            INTEGER     NOT NULL PRIMARY KEY AUTO_INCREMENT,
    id_user       INTEGER     NOT NULL,
    id_resolution INTEGER     NOT NULL,
    name          VARCHAR(30) NOT NULL,
    status_code   TINYINT     NOT NULL DEFAULT 0,
    created_at    DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at   DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE (id, id_user)
);

ALTER TABLE presentation
    ADD COLUMN id_resolution INTEGER NOT NULL;

CREATE TABLE IF NOT EXISTS page
(
    id              INTEGER  NOT NULL PRIMARY KEY AUTO_INCREMENT,
    id_presentation INTEGER  NOT NULL REFERENCES presentation (id),
    duration        INTEGER  NOT NULL,
    component       JSON     NOT NULL,
    status_code     TINYINT  NOT NULL DEFAULT 0,
    created_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at     DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS device_presentation
(
    id_device       INTEGER  NOT NULL REFERENCES device (id),
    id_presentation INTEGER  NOT NULL REFERENCES presentation (id),
    is_playing      INTEGER  NOT NULL DEFAULT 0,
    status_code     TINYINT  NOT NULL DEFAULT 0,
    created_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at     DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE (id_device, id_presentation)
);