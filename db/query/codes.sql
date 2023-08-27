-- name: CreateCode :one
INSERT INTO codes (
  code
) VALUES (
  $1
) RETURNING *;

-- name: UpdateCode :one
UPDATE codes SET
  phone = $1,
  email = $2,
  birthday = $3,
  name_surname = $4,
  address = $5,
  city_code = $6,
  registered_date = $7,
  ip = $8,
  active = $9,
  tc_no = $10,
  used = $11
WHERE code = $12
RETURNING *;

-- name: BanCode :one
UPDATE codes SET
  banned = $1,
  banned_reason = $2
WHERE code = $3
RETURNING *;

-- name: GetCode :one
SELECT * FROM codes WHERE code = $1;

-- name: GetUsedCodes :many
SELECT * FROM codes WHERE used = true and banned = false ORDER BY registered_date DESC;

-- name: DeleteCodes :exec
DELETE FROM codes;