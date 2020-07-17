CREATE TABLE IF NOT EXISTS state
(
    id        SMALLINT UNSIGNED AUTO_INCREMENT,
    status    TINYTEXT NOT NULL,
    info      TEXT,
    createdAt DATETIME DEFAULT current_timestamp,
    updatedAt DATETIME DEFAULT current_timestamp ON UPDATE current_timestamp,
    PRIMARY KEY (`id`)
);

INSERT INTO state (status, info) VALUES ('pause', '準備中...');

CREATE TABLE IF NOT EXISTS presentation
(
    id          SMALLINT UNSIGNED AUTO_INCREMENT,
    name        TEXT,
    speakers    TEXT,
    description TEXT,
    prev        SMALLINT,
    next        SMALLINT,
    createdAt   DATETIME DEFAULT current_timestamp,
    updatedAt   DATETIME DEFAULT current_timestamp ON UPDATE current_timestamp,
    PRIMARY KEY (`id`)
);

INSERT INTO presentation(id,name) VALUES (0,'before');

CREATE TABLE IF NOT EXISTS reaction
(
    id             SMALLINT UNSIGNED AUTO_INCREMENT,
    userId         VARCHAR(32)       NOT NULL,
    presentationId SMALLINT UNSIGNED NOT NULL,
    stamp          TINYTEXT,
    createdAt      DATETIME DEFAULT current_timestamp,
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS comment
(
    id             SMALLINT UNSIGNED AUTO_INCREMENT,
    userId         VARCHAR(32)       NOT NULL,
    presentationId SMALLINT UNSIGNED NOT NULL,
    text           TEXT,
    createdAt      DATETIME DEFAULT current_timestamp,
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS review
(
    userId         VARCHAR(32),
    presentationId SMALLINT UNSIGNED NOT NULL,
    skill          TINYINT UNSIGNED  NOT NULL,
    artistry       TINYINT UNSIGNED  NOT NULL,
    idea           TINYINT UNSIGNED  NOT NULL,
    presentation   TINYINT UNSIGNED  NOT NULL,
    createdAt      DATETIME DEFAULT current_timestamp,
    updatedAt      DATETIME DEFAULT current_timestamp ON UPDATE current_timestamp,
    PRIMARY KEY (`userId`, `presentationId`)
);
