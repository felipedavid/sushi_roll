CREATE TABLE games (
  id bigserial PRIMARY KEY,
  title varchar NOT NULL,
  description varchar NOT NULL,
  release_at timestamp NOT NULL,
  created_at timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE comments (
  id bigserial PRIMARY KEY,
  user_id bigint,
  game_id bigint,
  content varchar,
  created_at timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE users (
  id bigserial PRIMARY KEY,
  name varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  hashed_password varchar(60) NOT NULL,
  role_id bigint,
  created_at timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE roles (
  id bigserial PRIMARY KEY,
  name varchar NOT NULL
);

INSERT INTO roles (id, name) VALUES (0, 'admin');
INSERT INTO roles (id, name) VALUES (1, 'user');

ALTER TABLE comments ADD CONSTRAINT fk_comment_user
  FOREIGN KEY (user_id) REFERENCES users(id);
ALTER TABLE comments ADD CONSTRAINT fk_comment_game
  FOREIGN KEY (game_id) REFERENCES games(id);

ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);
ALTER TABLE users ADD CONSTRAINT fk_users_role
	FOREIGN KEY (role_id) REFERENCES roles(id);