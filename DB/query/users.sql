-- name: CreateUser :one
INSERT INTO users (
  name,username,age,email,password,address,contact,gender
) VALUES (
  $1, $2,$3,$4,$5,$6,$7,$8
)
RETURNING *;

-- name: LoginUser :one
SELECT * FROM users
WHERE email = $1
LIMIT 1;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;


-- name: UpdateUser :one
UPDATE users
SET name = $2,
    address = $3,
    contact = $4
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;