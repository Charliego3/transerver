-- name: AccountCreate :one
INSERT INTO accounts (
    create_at,
    user_id,
    username,
    region,
    area,
    phone,
    email,
    avatar,
    password,
    pwd_level,
    platform
) VALUES (
    NOW(), $1, $2, $3, $4,
    $5, $6, $7, $8, $9, $10
) RETURNING *;

-- name: AccountById :one
SELECT *
FROM accounts
WHERE id = $1;

-- name: AccountExistByPhone :one
SELECT true::bool
FROM accounts
WHERE phone = $1;

-- name: AccountExistsByEmail :one
SELECT true::bool
FROM accounts
WHERE email = $1;
