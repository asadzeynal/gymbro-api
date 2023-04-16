package repository

import (
	"context"
	"fmt"

	postgres "github.com/asadzeynal/gymbro-api/gen/sqlc"
	"github.com/asadzeynal/gymbro-api/internal/domain"
	"github.com/gofrs/uuid/v5"
	"github.com/jackc/pgx/v5"
)

type templateRepository struct {
	conn    *pgx.Conn
	queries postgres.Querier
}

func NewTemplateRepository(conn *pgx.Conn) *templateRepository {
	return &templateRepository{
		conn,
		postgres.New(conn),
	}
}

func (tr *templateRepository) execTx(ctx context.Context, fn func(queries *postgres.Queries) error) error {
	tx, err := tr.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	q := postgres.New(tx)
	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit(ctx)
}

func (tr *templateRepository) Store(ctx context.Context, template domain.Template) (uuid.UUID, error) {
	templateId, err := uuid.NewV4()
	if err != nil {
		return uuid.Nil, err
	}

	tr.execTx(ctx, func(q *postgres.Queries) error {
		_, err = q.AddTemplate(ctx, postgres.AddTemplateParams{
			ID:     templateId,
			Name:   template.Name,
			UserID: template.UserId,
		})

		exs := template.Exercises
		for i := range exs {
			exId, err := uuid.NewV4()
			if err != nil {
				return err
			}

			q.AddTemplateExercise(ctx, postgres.AddTemplateExerciseParams{
				ID:           exId,
				TemplateID:   templateId,
				ExerciseID:   exs[i].ExerciseId,
				DisplayOrder: exs[i].DisplayOrder,
			})

			sets := exs[i].Sets
			for j := range sets {
				setId, err := uuid.NewV4()
				if err != nil {
					return err
				}

				q.AddTemplateSet(ctx, postgres.AddTemplateSetParams{
					ID:                 setId,
					TemplateExerciseID: exId,
					Reps:               sets[j].Reps,
					Weight:             float64(sets[j].Weight),
				})
			}
		}
		return nil
	})
	return templateId, nil
}
