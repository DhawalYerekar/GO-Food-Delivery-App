-- name: CreateHotel :one
INSERT INTO hotel (
  name,email,address,contact
) VALUES (
  $1, $2,$3,$4
)
RETURNING *;

-- name: GetHotel :one
SELECT * FROM hotel
WHERE id = $1 LIMIT 1;


-- name: UpdateHotel :one
UPDATE hotel
SET name = $2,
    address = $3,
    contact = $4
WHERE id = $1
RETURNING *;

-- name: DeleteHotel :exec
DELETE FROM hotel
WHERE id = $1;

-- -- name: GetHotelsByDishName :many
-- SELECT DISTINCT h.*
-- FROM hotel h
-- JOIN dishes d ON h.id = d.hotel_id
-- WHERE d.name = $1;