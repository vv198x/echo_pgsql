DROP TABLE IF EXISTS users CASCADE;

CREATE TABLE users
(
    login VARCHAR(20) NOT NULL UNIQUE PRIMARY KEY,
    password VARCHAR(20) NOT NULL,
    rule INTEGER,
    name VARCHAR(30) NOT NULL,
    last_name VARCHAR(30) NOT NULL,
    dob INTEGER
);

INSERT INTO users
VALUES ('admin', 'admin', 0, 'first name Admin', 'last name Admin', 0)
RETURNING *;

INSERT INTO users
VALUES ('user', 'user', 1, 'first name user', 'last name user', 1000000)
RETURNING *;

INSERT INTO users
VALUES ('lock', 'lock', 2, 'first name lock', 'last name lock', 2000000)
RETURNING *;