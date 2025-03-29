-- name: CreateProduct :one
INSERT INTO products ("id", "seller_id", "product_name", "description", "base_price", "auction_end")
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id;

-- name: GetProductById :one
SELECT * FROM products
WHERE id = $1;