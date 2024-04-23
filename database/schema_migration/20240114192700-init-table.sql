
-- +migrate Up

CREATE TABLE products (
	id serial4 NOT NULL,
	"name" text NOT NULL,
	price float8 NOT NULL,
	stock int8 NOT NULL,
	description text NOT NULL,
	image_url text NOT NULL,
	created_at timestamp NOT NULL DEFAULT '2024-03-24 01:46:25.635215'::timestamp without time zone,
	updated_at timestamp NOT NULL,
	deleted_at timestamp NOT NULL,
	CONSTRAINT products_pkey PRIMARY KEY (id)
);

-- +migrate Down
