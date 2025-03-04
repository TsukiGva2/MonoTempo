PRAGMA synchronous = OFF;
PRAGMA journal_mode = MEMORY;
BEGIN TRANSACTION;

CREATE TABLE `athletes` (
  `num`      INTEGER NOT NULL PRIMARY KEY
,  `event_id` integer DEFAULT NULL
,  `name`     TEXT
,  `city`     TEXT
,  `team`     TEXT
,  `track_id` integer DEFAULT NULL
,  `sex`      TEXT
);
CREATE TABLE `equipamento` (
  `id`       INTEGER NOT NULL PRIMARY KEY
,  `idequip`  INTEGER NOT NULL
,  `modelo`   TEXT    NOT NULL
,  `event_id` integer NOT NULL
);
CREATE TABLE `event_data` (
  `id`          INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT
,  `event_date`  TEXT DEFAULT NULL
,  `event_title` TEXT
);
CREATE TABLE `staffs` (
  `id`       INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT
,  `event_id` integer NOT NULL
,  `nome`     TEXT DEFAULT NULL
);
CREATE TABLE `tracks` (
  `id`               INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT
,  `event_id`         INTEGER DEFAULT NULL
,  `race_description` TEXT
,  `inicio`           TEXT DEFAULT NULL
,  `chegada`          TEXT DEFAULT NULL
,  `largada`          TEXT DEFAULT NULL
);

INSERT INTO equipamento(id, idequip, modelo, event_id) VALUES (1,0,'',0); 
END TRANSACTION;
