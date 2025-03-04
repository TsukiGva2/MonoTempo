PRAGMA synchronous = OFF;
PRAGMA journal_mode = MEMORY;
BEGIN TRANSACTION;

CREATE TABLE athletes
(
   num              INTEGER NOT NULL PRIMARY KEY
,  event_id         INTEGER NOT NULL
,  track_id         INTEGER NOT NULL
,  name             TEXT
,  city             TEXT
,  team             TEXT
,  sex              TEXT
);
CREATE TABLE equipamento
(
   id               INTEGER NOT NULL PRIMARY KEY
,  idequip          INTEGER NOT NULL
,  event_id         INTEGER NOT NULL
,  modelo           TEXT    NOT NULL
);
CREATE TABLE event_data
(
   id               INTEGER NOT NULL PRIMARY KEY
,  event_date       TEXT DEFAULT NULL
,  event_title      TEXT
);
CREATE TABLE tracks
(
   id               INTEGER NOT NULL PRIMARY KEY
,  event_id         INTEGER DEFAULT NULL
,  inicio           TEXT DEFAULT NULL
,  chegada          TEXT DEFAULT NULL
,  largada          TEXT DEFAULT NULL
,  race_description TEXT
);
CREATE TABLE staffs
(
   id               INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT
,  event_id         INTEGER NOT NULL
,  nome             TEXT DEFAULT NULL
); 

INSERT INTO equipamento(id, idequip, modelo, event_id) VALUES (1,0,'',0); 
END TRANSACTION;
