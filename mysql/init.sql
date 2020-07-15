CREATE TABLE IF NOT EXISTS state
(
    id        SMALLINT AUTO_INCREMENT,
    state     TINYTEXT NOT NULL DEFAULT 'pause',
    info      TEXT              DEFAULT '準備中...',
    createdAt DATETIME,
    updatedAt DATETIME,
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS presentation
(
    id          SMALLINT AUTO_INCREMENT,
    name        TEXT,
    speakers    TEXT,
    description TEXT,
    prev        SMALLINT,
    next        SMALLINT,
    createdAt   DATETIME,
    updatedAt   DATETIME,
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS reaction
(
    id             SMALLINT AUTO_INCREMENT,
    userId         VARCHAR(32)      NOT NULL,
    presentationId TINYINT UNSIGNED NOT NULL,
    stamp          TINYTEXT,
    createdAt      DATETIME,
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS comment
(
    id             SMALLINT AUTO_INCREMENT,
    userId         VARCHAR(32)      NOT NULL,
    presentationId TINYINT UNSIGNED NOT NULL,
    text           TEXT,
    createdAt      DATETIME,
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS review
(
    userId         VARCHAR(32),
    presentationId SMALLINT UNSIGNED,
    skill          TINYINT UNSIGNED NOT NULL,
    artistry       TINYINT UNSIGNED NOT NULL,
    idea           TINYINT UNSIGNED NOT NULL,
    presentation   TINYINT UNSIGNED NOT NULL,
    createdAt      DATETIME,
    updatedAt      DATETIME,
    PRIMARY KEY (`userId`, `presentationId`)
);