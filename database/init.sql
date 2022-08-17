CREATE TABLE IF NOT EXISTS orders
(
    id serial not null primary key,
    info jsonb not null
)
