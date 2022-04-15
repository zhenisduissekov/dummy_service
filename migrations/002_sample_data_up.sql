-- +goose Up
INSERT INTO orders (abc_order_id, def_id, error_code, error_message, action, order_status)
VALUES('123x123x123', 'xxx77er', 1, 'balance too low', 11, 'error'),
    ('123x123x123', 'zzz77er', 2, 'authorization failed', 12, 'error'),
    ('123x444x123', 'cccc77er', 3, 'timeout', 11, 'error'),
    ('123x444x123', 'sss77er', 4, 'connection closed', 11, 'error'),
    ('123x444x555', 'avx77er', 5, 'balance too low', 11, 'error')
;

-- +goose Down
TRUNCATE TABLE orders;

