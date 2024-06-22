CREATE DATABASE kakebo_api;

CREATE TABLE groups (
  id bigserial PRIMARY KEY,
  revision bigint DEFAULT 0,
  updated_at TIMESTAMP NOT NULL,
  created_at TIMESTAMP NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE users (
  id bigserial PRIMARY KEY,
  uid varchar(255) NOT NULL,
  group_id BIGINT,
  group_admin int NOT NULL DEFAULT 0,
  name varchar(255) DEFAULT NULL,
  register_type int DEFAULT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  UNIQUE (email),
  FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE RESTRICT
);

CREATE TABLE events (
  id bigserial PRIMARY KEY,
  amount int NOT NULL,
  category int DEFAULT NULL,
  date varchar(255) DEFAULT NULL,
  store_name varchar(255) DEFAULT NULL,
  group_id bigint NOT NULL,
  memo varchar(255) DEFAULT NULL,
  create_user varchar(255) DEFAULT NULL,
  update_user varchar(255) DEFAULT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE RESTRICT
);

CREATE TABLE privates (
  id bigserial PRIMARY KEY,
  amount int NOT NULL,
  category int DEFAULT NULL,
  date varchar(255) DEFAULT NULL,
  store_name varchar(255) DEFAULT NULL,
  user_id bigint NOT NULL,
  memo varchar(255) DEFAULT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE RESTRICT
);

CREATE TABLE patterns (
  id serial PRIMARY KEY,
  user_id bigint NOT NULL,
  store_name varchar(255) NOT NULL,
  category int NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  created_at TIMESTAMP NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE RESTRICT
);
