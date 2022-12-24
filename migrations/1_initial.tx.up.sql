CREATE TABLE users
(
    login VARCHAR(20) NOT NULL UNIQUE PRIMARY KEY,
    password VARCHAR(60) NOT NULL,
    rule INTEGER,
    name VARCHAR(30) NOT NULL,
    last_name VARCHAR(30) NOT NULL,
    dob INTEGER
);

INSERT INTO users
VALUES ('admin', '$2a$10$oTstAZJ1Uj2mO4brcs5yke9i0WHg0nNf9rVHQs55Y4ayNheOf1NzC', 0, 'first name Admin', 'last name Admin', 0)
RETURNING *;

INSERT INTO users
VALUES ('user', '$2a$10$z1co3zjNjhBHftOeqelx3e20yQKFCCMxkF6OS9JCCkzFlfYrxeRoG', 1, 'first name user', 'last name user', 100000000)
RETURNING *;

INSERT INTO users
VALUES ('lock', '$2a$10$W4Sky819LgwPaK7ABzW/L.86VXsp8VlLyjciiaolnmkTCcxoJjlIm', 2, 'first name lock', 'last name lock', 200000000)
RETURNING *;