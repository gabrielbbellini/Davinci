CREATE TABLE IF NOT EXISTS presentation
(
    id             INTEGER     NOT NULL PRIMARY KEY AUTO_INCREMENT,
    id_resolution  INTEGER     NOT NULL REFERENCES resolution (id),
    id_orientation INTEGER     NOT NULL,
    id_user        INTEGER     NOT NULL REFERENCES user (id),
    name           VARCHAR(30) NOT NULL,
    status_code    TINYINT     NOT NULL DEFAULT 0,
    created_at     DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at    DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

ALTER TABLE presentation
    ADD CONSTRAINT presentation_resolution_fk
        FOREIGN KEY (id_resolution) REFERENCES resolution (id);


ALTER TABLE presentation
    ADD CONSTRAINT presentation_user_fk
        FOREIGN KEY (id_user) REFERENCES user (id);

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

ALTER TABLE page
    ADD CONSTRAINT page_presentation_fk
        FOREIGN KEY (id_presentation) REFERENCES presentation (id);

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

ALTER TABLE device_presentation
    ADD CONSTRAINT device_presentation_presentation
        FOREIGN KEY (id_presentation) REFERENCES presentation (id);

ALTER TABLE device_presentation
    ADD CONSTRAINT device_presentation_device
        FOREIGN KEY (id_device) REFERENCES device (id);