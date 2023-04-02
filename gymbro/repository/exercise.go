package repository

import (
	"context"
	"database/sql"

	"github.com/asadzeynal/gymbro-api/domain"
	postgres "github.com/asadzeynal/gymbro-api/gen/sqlc"
	"github.com/google/uuid"
)

type exerciseRepository struct {
	conn    *sql.DB
	queries postgres.Querier
}

func NewExerciseRepository(conn *sql.DB) *exerciseRepository {
	return &exerciseRepository{
		conn,
		postgres.New(conn),
	}
}

func (e *exerciseRepository) Store(ctx context.Context, exercise *domain.Exercise) (uuid.UUID, error) {
	uuid, err := e.queries.AddExercise(ctx, postgres.AddExerciseParams{
		ID:   exercise.ID,
		Name: exercise.Name,
		Description: sql.NullString{
			String: exercise.Description,
			Valid:  true,
		},
	})
	return uuid, err
}

func (e *exerciseRepository) Delete(ctx context.Context, id uuid.UUID) error {
	err := e.queries.DeleteExercise(ctx, id)
	return err
}

func (e *exerciseRepository) Fetch(ctx context.Context) ([]domain.Exercise, error) {
	excs, err := e.queries.FetchExercises(ctx)
	if err != nil {
		return []domain.Exercise{}, err
	}

	result := make([]domain.Exercise, len(excs))
	for i := 0; i < len(excs); i++ {
		ex := excs[i]
		result[i] = convertToDomain(ex)
	}

	return result, nil
}

func (e *exerciseRepository) GetById(ctx context.Context, id uuid.UUID) (domain.Exercise, error) {
	ex, err := e.queries.GetExerciseById(ctx, id)
	if err != nil {
		return domain.Exercise{}, err
	}

	return convertToDomain(ex), nil
}

func convertToDomain(dbEx postgres.Exercise) domain.Exercise {
	return domain.Exercise{
		ID:          dbEx.ID,
		Name:        dbEx.Name,
		Description: dbEx.Description.String,
	}
}
