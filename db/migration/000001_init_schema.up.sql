CREATE TABLE game (
  id bigserial PRIMARY KEY,
  title varchar NOT NULL,
  description varchar NOT NULL,
  release_at timestamp NOT NULL,
  created_at timestamp NOT NULL DEFAULT (now())
);
