
-- +migrate Up
-- products definition

-- Drop table

-- DROP TABLE products;

CREATE TABLE products (
	id BIGSERIAL NOT NULL,
	"name" text NOT NULL,
	price float8 NOT NULL,
	stock int8 NOT NULL,
	description text NOT NULL,
	image_url text NOT NULL,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT products_pkey PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE products;