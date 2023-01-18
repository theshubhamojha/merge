CREATE TABLE IF NOT EXISTS configurations (
  id varchar primary key,
  "configuration" json,
  "role" role_type,
  created_at timestamp,
  updated_at timestamp
);