CREATE TABLE game (
  id bigserial PRIMARY KEY,
  title varchar NOT NULL,
  description varchar NOT NULL,
  release timestamp NOT NULL,
  create_at timestamp NOT NULL DEFAULT (now())
);
