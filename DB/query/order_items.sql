-- name: CreateOrderItem :one
INSERT INTO order_items (
  order_id,food_id,quantity
) VALUES (
  $1, $2,$3
)
RETURNING *;

-- name: GetOrderItem :one
SELECT * FROM order_items
WHERE id = $1 LIMIT 1;

-- name: ListOrderItems :many
SELECT * FROM order_items
WHERE order_id = $1;

-- name: UpdateOrderItem :one
UPDATE order_items
SET food_id = $2
WHERE id = $1
RETURNING *;

-- name: UpdateOrderItemQuantity :one
UPDATE order_items
SET quantity = $2
WHERE id = $1
RETURNING *;

-- name: DeleteOrderItem :exec
DELETE FROM order_items
WHERE id = $1;