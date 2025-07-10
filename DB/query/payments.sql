-- name: CreatePayment :one
INSERT INTO payments (
  order_id,payment_mode,status
) VALUES (
  $1, $2,$3
)
RETURNING *;

-- name: GetPayment :one
SELECT * FROM payments
WHERE id = $1 LIMIT 1;

-- name: GetPaymentsByOrderID :many
SELECT * FROM payments
WHERE order_id = $1;

-- -- name: ListOrdersByUserID :many
-- SELECT * FROM orders
-- WHERE user_id = $1
-- ORDER BY order_time DESC;

-- name: UpdatePaymentStatus :one
UPDATE payments
SET status = $2
WHERE id = $1
RETURNING *;

-- name: UpdatePaymentMode :one
UPDATE payments
SET payment_mode = $2
WHERE id = $1
RETURNING *;

-- name: DeletePayment :exec
DELETE FROM payments
WHERE id = $1;