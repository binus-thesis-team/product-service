
-- +migrate Up
-- products definition

-- Drop table

-- DROP TABLE products;

CREATE TABLE products (
	id serial4 NOT NULL,
	"name" text NOT NULL,
	price float8 NOT NULL,
	stock int8 NOT NULL,
	description text NOT NULL,
	image_url text NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	deleted_at timestamp NULL,
	CONSTRAINT products_pkey PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE products;