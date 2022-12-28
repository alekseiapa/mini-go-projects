-- name: CreateBook :one
INSERT INTO book (
    name,
    publication
)
VALUES (
    $1, $2
)
RETURNING *;

-- name: GetBook :one
SELECT * FROM book
WHERE uuid = $1 LIMIT 1;


-- name: ListBooks :many
SELECT * FROM book
ORDER BY publication
LIMIT $1
OFFSET $2;

-- name: UpdateBookName :one
UPDATE book
    SET name = $2
WHERE uuid = $1
RETURNING *;


-- name: UpdateBookPublication :one
UPDATE book
    SET publication = $2
WHERE uuid = $1
RETURNING *;

-- name: DeleteBook :exec
DELETE FROM book
WHERE uuid = $1;