-- name: RegionList :many
SELECT * FROM regions;


-- name: RegionByCode :one
SELECT * FROM regions WHERE code = $1;


-- name: RegionCreate :one
INSERT INTO regions(code, area, img, name) VALUES ($1, $2, $3, $4) RETURNING *;
