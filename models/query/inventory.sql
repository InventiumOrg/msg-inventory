-- name: CreateInventory :one
INSERT INTO inventory (
    name, quantity, category, located
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetInventory :one
SELECT * FROM inventory 
WHERE id = $1 LIMIT 1;

-- -- name: GetInventoryByPOS :one
-- SELECT * FROM inventory 
-- WHERE pos_id = $1 LIMIT 1;

-- name: ListInventories :many
SELECT * FROM inventory
ORDER BY id
LIMIT $1
OFFSET $2;

-- -- name: ListInventoriesByPOS :many
-- SELECT * FROM inventory
-- WHERE pos_id = $2
-- ORDER BY id
-- LIMIT $1
-- OFFSET $3;

-- name: UpdateInventory :one
UPDATE inventory
SET quantity = $2
WHERE id = $1
RETURNING *;

-- -- name: UpdateInventoryPOS :one
-- UPDATE inventory
-- SET quantity = $3
-- WHERE pos_id = $1
-- AND id = $1
-- RETURNING *;

-- name: DeleteInventory :exec
DELETE FROM inventory
WHERE id = $1;