-- +goose Up
CREATE TABLE products (
  id          BIGSERIAL      PRIMARY KEY,
  name        VARCHAR(30)    NOT NULL,
  price       NUMERIC(10,2)  NOT NULL,
  created_at  TIMESTAMPTZ    NOT NULL DEFAULT clock_timestamp(),
  updated_at  TIMESTAMPTZ    NOT NULL DEFAULT clock_timestamp()
);

CREATE UNIQUE INDEX products_name_uniq_idx ON products(LOWER(name));

-- +goose Down
DROP INDEX IF EXISTS products_name_uniq_idx;
DROP TABLE users;
