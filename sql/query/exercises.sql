-- name: AddExercise :one
INSERT INTO exercises (id, name, description)
    VALUES ($1, $2, $3)
RETURNING
    id;

-- name: FetchExercises :many
SELECT
    id,
    name,
    description
FROM
    exercises
ORDER BY
    name ASC;

-- name: DeleteExercise :exec
DELETE FROM exercises
WHERE id = $1;

-- name: GetExerciseById :one
SELECT
    id,
    name,
    description
FROM
    exercises
WHERE
    id = $1;

