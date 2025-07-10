-- name: CreateDish :one
INSERT INTO dishes (
  name,type,price,hotel_id
) VALUES (
  $1, $2,$3,$4
)
RETURNING *;

-- name: GetDish :one
SELECT * FROM dishes
WHERE id = $1 LIMIT 1;

-- name: GetDishesByHotelID :many
SELECT * FROM dishes
WHERE hotel_id = $1;

-- name: ListDishesByName :many
SELECT * FROM dishes
WHERE name = $1;

-- name: UpdateDish :one
UPDATE dishes
SET name = $2,
    type = $3,
    price = $4
WHERE id = $1
RETURNING *;

-- name: DeleteDish :exec
DELETE FROM dishes
WHERE id = $1;