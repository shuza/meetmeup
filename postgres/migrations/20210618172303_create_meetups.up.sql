CREATE TABLE meetups
(
    id          BIGSERIAL PRIMARY KEY,
    name        VARCHAR(255) UNIQUE                               NOT NULL,
    description TEXT,

    user_id     BIGSERIAL REFERENCES users (id) ON DELETE CASCADE NOT NULL
);