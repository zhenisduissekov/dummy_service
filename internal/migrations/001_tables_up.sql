-- +goose Up
CREATE TABLE orders (
                             id serial PRIMARY KEY,
                             abc_order_id VARCHAR(255),
                             def_id VARCHAR(255),
                             error_code INT,
                             error_message VARCHAR(255),
                             action INT,
                             order_status VARCHAR(255),
                             created_at timestamp default now()
);

-- +goose Down
DROP TABLE orders;

