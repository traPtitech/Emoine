CREATE TABLE IF NOT EXISTS state
(
    id        SMALLINT AUTO_INCREMENT,
    state     TINYTEXT NOT NULL,
    info      TEXT,
    createdAt DATETIME NOT NULL,
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
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS reaction
(
    id        SMALLINT AUTO_INCREMENT,
    userId    VARCHAR(32)      NOT NULL,
    team      TINYINT UNSIGNED NOT NULL,
    stamp     TINYTEXT,
    comment   TEXT,
    createdAt DATETIME,
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS review
(
    userId       VARCHAR(32),
    teamId       SMALLINT UNSIGNED,
    skill        TINYINT UNSIGNED NOT NULL,
    artistry     TINYINT UNSIGNED NOT NULL,
    idea         TINYINT UNSIGNED NOT NULL,
    presentation TINYINT UNSIGNED NOT NULL,
    createdAt    DATETIME,
    updatedAt    DATETIME,
    PRIMARY KEY (userId, teamId)
);