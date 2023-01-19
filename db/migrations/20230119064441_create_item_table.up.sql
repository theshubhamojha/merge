CREATE TABLE IF NOT EXISTS items (
  id varchar primary key,
  account_id varchar,
  "name" varchar,
  image_url varchar,
  quantity int,
  is_active boolean,
  created_at timestamp,
  updated_at timestamp
);

ALTER TABLE items
ADD FOREIGN KEY (account_id) REFERENCES accounts(id);