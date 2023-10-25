CREATE TABLE users (
    id         bigserial PRIMARY KEY,
    name       text NOT NULL,
    email      text UNIQUE NOT NULL,
    password   text NOT NULL
);
