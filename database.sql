-- create value status as enum
CREATE TYPE status AS ENUM ('active', 'inactive');

-- create table users with relation customers
CREATE TABLE users (
    id serial not null,
    username varchar (20) not null,
    password varchar (100) not null,
    role varchar (20) not null,
    email varchar (50) not null,
    email_verify status default 'inactive',
    -- customer_id integer references customers(customer_id), -- relation from customres table
    created_at integer not null,
    CONSTRAINT users_pkey PRIMARY KEY (id)
);

-- give check for username and password not allowed empty string
ALTER TABLE users ADD CONSTRAINT username_check check (username != '');
ALTER TABLE users ADD CONSTRAINT password_check check (password != '');

CREATE TABLE user_token (
    id serial not null,
    email varchar(50) not null,
    token varchar(255) not null,
    date_created integer,
    CONSTRAINT token_pkey PRIMARY KEY (id)
);

-- -- create table customers
-- CREATE TABLE customers (
--     customer_id serial primary key,
--     name varchar (25) not null,
--     date_of_birth date not null,
--     zip_code varchar(10) not null,
--     status status default 'inactive',
--     created_at timestamp default now()
-- );

-- -- create index for searching with where clause
-- CREATE INDEX customers_zip_code_and_status ON customers(zip_code, status);



-- -- create table category
-- CREATE TABLE categories (
--     category_id serial not null,
--     category_name varchar(20) not null,
--     CONSTRAINT category_pkey PRIMARY KEY (category_id)
-- );

-- -- create table product
-- CREATE TABLE products (
--     product_id serial not null,
--     product_name varchar(255) not null,
--     category_id integer references categories(category_id),
--     price decimal DEFAULT 0,
--     stock integer DEFAULT 0,
--     product_description text,
--     CONSTRAINT product_pkey PRIMARY KEY (product_id)
-- );

-- -- create index for searching where cluase by category and stock
-- CREATE INDEX product_category_status ON products(category_id,stock);

-- ALTER TABLE products ADD CONSTRAINT price_check check (price >= 0);
-- ALTER TABLE products ADD CONSTRAINT stock_check check (stock >= 0);

-- -- create table images
-- CREATE TABLE images (
--     product_id integer references products(product_id),
--     image_url varchar(255)
-- );

