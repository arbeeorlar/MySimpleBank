-- name: CreateAccount :one
INSERT INTO account (
    owner, balance,currency
) VALUES (
             $1, $2,$3
         )
    RETURNING *;

-- name: GetAuthor :one
SELECT * FROM authors
WHERE id = $1 LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM authors
ORDER BY name;
