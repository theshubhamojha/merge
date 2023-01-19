CREATE TABLE IF NOT EXISTS cart (
  id varchar primary key,
  account_id varchar,
  item_id varchar,
  quantity int, 
  is_deleted boolean,
  created_at timestamp,
  updated_at timestamp
);

ALTER TABLE cart
ADD FOREIGN KEY (account_id) REFERENCES accounts(id);

ALTER TABLE cart
ADD FOREIGN KEY (item_id) REFERENCES items(id);