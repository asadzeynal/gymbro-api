// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: exercises.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const addExercise = `-- name: AddExercise :one
INSERT INTO exercises (id, name, description)
    VALUES ($1, $2, $3)
RETURNING
    id
`

type AddExerciseParams struct {
	ID          uuid.UUID
	Name        string
	Description string
}

func (q *Queries) AddExercise(ctx context.Context, arg AddExerciseParams) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, addExercise, arg.ID, arg.Name, arg.Description)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const deleteExercise = `-- name: DeleteExercise :exec
DELETE FROM exercises
WHERE id = $1
`

func (q *Queries) DeleteExercise(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteExercise, id)
	return err
}

const fetchExercises = `-- name: FetchExercises :many
SELECT
    id,
    name,
    description
FROM
    exercises
ORDER BY
    name ASC
`

func (q *Queries) FetchExercises(ctx context.Context) ([]Exercise, error) {
	rows, err := q.db.QueryContext(ctx, fetchExercises)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Exercise
	for rows.Next() {
		var i Exercise
		if err := rows.Scan(&i.ID, &i.Name, &i.Description); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getExerciseById = `-- name: GetExerciseById :one
SELECT
    id,
    name,
    description
FROM
    exercises
WHERE
    id = $1
`

func (q *Queries) GetExerciseById(ctx context.Context, id uuid.UUID) (Exercise, error) {
	row := q.db.QueryRowContext(ctx, getExerciseById, id)
	var i Exercise
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}
