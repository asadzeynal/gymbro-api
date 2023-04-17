package repository

import (
	"context"
	"database/sql"

	postgres "github.com/asadzeynal/gymbro-api/gen/sqlc"
	"github.com/asadzeynal/gymbro-api/internal/domain"
	"github.com/gofrs/uuid/v5"
	"github.com/jackc/pgx/v5"
)

type exerciseRepository struct {
	conn    *pgx.Conn
	queries postgres.Querier
}

func NewExerciseRepository(conn *pgx.Conn) *exerciseRepository {
	return &exerciseRepository{
		conn,
		postgres.New(conn),
	}
}

func (e *exerciseRepository) Store(ctx context.Context, exercise *domain.Exercise) (uuid.UUID, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return uuid.UUID{}, err
	}
	uuid, err := e.queries.AddExercise(ctx, postgres.AddExerciseParams{
		ID:          id,
		Name:        exercise.Name,
		Description: exercise.Description,
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
		if err == sql.ErrNoRows {
			return []domain.Exercise{}, nil
		}
		return []domain.Exercise{}, err
	}

	result := make([]domain.Exercise, len(excs))
	for i := 0; i < len(excs); i++ {
		ex := excs[i]
		result[i] = exerciseToDomain(ex)
	}

	return result, nil
}

func (e *exerciseRepository) GetById(ctx context.Context, id uuid.UUID) (domain.Exercise, error) {
	ex, err := e.queries.GetExerciseById(ctx, id)
	if err != nil {
		return domain.Exercise{}, err
	}

	return exerciseToDomain(ex), nil
}
