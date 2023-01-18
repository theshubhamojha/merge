CREATE TYPE role_type AS ENUM ('admin', 'user');

CREATE TABLE IF NOT EXISTS accounts (
  id varchar primary key,
  "name" varchar,
  email varchar,
  "password" varchar,
  "role" role_type,
  is_active boolean,
  created_at timestamp,
  updated_at timestamp
);