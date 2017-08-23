CREATE TABLE lookup(
   id             SERIAL PRIMARY KEY,
   isbn           TEXT    NOT NULL,
   results        INT     NOT NULL
);