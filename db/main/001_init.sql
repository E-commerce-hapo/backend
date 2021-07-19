CREATE TABLE IF NOT EXISTS category (
    id int8 NOT NULL PRIMARY KEY UNIQUE,
    name VARCHAR,
    description VARCHAR,
    shop_id int8,
    created_at TIMESTAMP,
    deleted_at TIMESTAMP,
    updated_at TIMESTAMP
)