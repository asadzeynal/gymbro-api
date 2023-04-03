package domain

import (
	"context"

	"github.com/google/uuid"
)

// Exercise is representing the Exercise data struct
type Exercise struct {
	ID          uuid.UUID
	Name        string
	Description string
}

// ExerciseRepository represents the exercise's repository contract
type ExerciseRepository interface {
	Store(ctx context.Context, exercise *Exercise) (uuid.UUID, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Fetch(ctx context.Context) ([]Exercise, error)
	GetById(ctx context.Context, id uuid.UUID) (Exercise, error)
}


