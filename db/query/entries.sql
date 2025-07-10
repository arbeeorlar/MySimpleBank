
-- "account_id" bigint NOT NULL,
--                            "amount" bigint NOT NULL,
-- name: CreateEntries :one
INSERT INTO entries (
    account_id, amount
) VALUES (
             $1, $2
         )
    RETURNING *;

-- name: GetEntries :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: ListEntriesByAccountId :many
SELECT * FROM entries
WHERE account_id=$1
ORDER BY id ;


-- name: UpdateAccount :one
-- UPDATE entries
-- set owner = $2
-- WHERE id = $1
--     RETURNING *;
--
-- -- name: DeleteAccount :exec
-- DELETE FROM account
-- WHERE id = $1;
