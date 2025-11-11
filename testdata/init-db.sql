CREATE TABLE player
(
    id     SERIAL PRIMARY KEY,
    health INTEGER NOT NULL,
    level  INTEGER NOT NULL
);

INSERT INTO player (health, level)
VALUES (100, 1),
       (150, 5);