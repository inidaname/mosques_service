-- name: CreateMosque :one
INSERT INTO mosques (id, name, address, eid_time, jummah_time, lat, lng, created_at)
VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5, $6, NOW())
RETURNING *;

-- name: GetUMosqueByName :one
SELECT * FROM mosques
WHERE name = $1 LIMIT 1;

-- name: GetAllMosque :many
SELECT * FROM mosques;

-- name: UpdateMosque :one
UPDATE mosques
SET
  name = coalesce(sqlc.narg(name), name),
  address = coalesce(sqlc.narg(address), address),
  eid_time = coalesce(sqlc.narg(eid_time), eid_time),
  jummah_time = coalesce(sqlc.narg(jummah_time), jummah_time),
  lat = coalesce(sqlc.narg(lat), lat),
  lng = coalesce(sqlc.narg(lng), lng)
WHERE id = sqlc.arg(id)
RETURNING *;