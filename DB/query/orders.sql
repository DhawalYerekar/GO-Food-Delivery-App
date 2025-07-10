-- name: CreateOrder :one
INSERT INTO orders (
  user_id,hotel_id,status,total_amount
) VALUES (
  $1, $2,$3,$4
)
RETURNING *;

-- name: GetOrder :one
SELECT * FROM orders
WHERE id = $1 LIMIT 1;

-- name: ListOrdersByUserID :many
SELECT * FROM orders
WHERE user_id = $1
ORDER BY order_time DESC;

-- name: UpdateOrderStatus :one
UPDATE orders
SET status = $2
WHERE id = $1
RETURNING *;

-- name: UpdateOrderAmount :one
UPDATE orders
SET total_amount = $2
WHERE id = $1
RETURNING *;

-- name: DeleteOrder :exec
DELETE FROM orders
WHERE id = $1;