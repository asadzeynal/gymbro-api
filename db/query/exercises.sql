-- name: AddExercise :one
INSERT INTO exercises (id, name, description)
    VALUES ($1, $2, $3)
RETURNING
    id;

