package service

import (
	"context"
	"time"

	"github.com/asadzeynal/gymbro-api/internal/domain"
)

type exerciseService struct {
	exerciseRepo   domain.ExerciseRepository
	contextTimeout time.Duration
}

func NewExerciseService(er domain.ExerciseRepository, timeout time.Duration) *exerciseService {
	return &exerciseService{er, timeout}
}

// Returns Exercises grouped by first letter
func (es *exerciseService) FetchGroupedByInitial(c context.Context) (map[string][]domain.Exercise, error) {
	ctx, cancel := context.WithTimeout(c, es.contextTimeout)
	defer cancel()

	res, err := es.exerciseRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}
	grouped := groupExerciseByInitial(res)
	return grouped, nil
}

func groupExerciseByInitial(exs []domain.Exercise) map[string][]domain.Exercise {
	grouped := make(map[string][]domain.Exercise)
	for i := range exs {
		initial := exs[i].Name[:1]
		grouped[initial] = append(grouped[initial], exs[i])
	}
	return grouped
}
