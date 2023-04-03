-- name: AddTemplate :one
INSERT INTO templates (id, name, user_id)
    VALUES ($1, $2, $3)
RETURNING
    id;

