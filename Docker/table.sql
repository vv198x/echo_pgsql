DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS users_db CASCADE;

CREATE TABLE users
(
    login VARCHAR(20) NOT NULL UNIQUE PRIMARY KEY,
    password VARCHAR(20) NOT NULL,
    rule INTEGER,
    name VARCHAR(30) NOT NULL,
    last_name VARCHAR(30) NOT NULL,
    dob DATE NOT NULL
);

INSERT INTO users
VALUES ('admin', 'admin', 0, 'first name Admin', 'last name Admin', '01.01.1970')
RETURNING *;

INSERT INTO users
VALUES ('user', 'user', 1, 'first name user', 'last name user', '11.11.1997')
RETURNING *;

INSERT INTO users
VALUES ('lock', 'lock', 2, 'first name lock', 'last name lock', '12.21.2000')
RETURNING *;