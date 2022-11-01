CREATE TABLE games (
  id bigserial PRIMARY KEY,
  title varchar NOT NULL,
  description varchar NOT NULL,
  release_at timestamp NOT NULL,
  created_at timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE users (
    id bigserial PRIMARY KEY,
    name varchar not null,
    email varchar not null,
    birth timestamp NOT NULL,
    hashed_password CHAR(60) not null,
    created_at timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);