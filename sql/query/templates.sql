-- name: AddTemplateSet :exec
INSERT INTO template_sets(id, reps, weight, template_exercise_id)
    VALUES ($1, $2, $3, $4);

-- name: AddTemplateExercise :exec
INSERT INTO template_exercises(id, template_id, exercise_id, display_order)
    VALUES ($1, $2, $3, $4);

-- name: AddTemplate :one
INSERT INTO templates(id, name, user_id)
    VALUES ($1, $2, $3)
RETURNING
    id;

-- name: GetTemplatesByUserId :many
SELECT
    *
FROM
    templates
WHERE
    user_id = $1;

-- name: GetTemplateExercisesByTemplateIds :many
SELECT
    *
FROM
    template_exercises
WHERE
    template_id IN ($1::uuid[]);

-- name: GetTemplateSetsByTemplateExerciseIds :many
SELECT
    *
FROM
    template_sets
WHERE
    template_exercise_id IN ($1::uuid[]);

