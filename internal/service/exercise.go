package service

import (
	"context"
	"fmt"
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

func (es *exerciseService) Store(c context.Context, ex *domain.Exercise) (string, error) {
	ctx, cancel := context.WithTimeout(c, es.contextTimeout)
	defer cancel()

	id, err := es.exerciseRepo.Store(ctx, ex)
	if err != nil {
		return "", fmt.Errorf("error when saving the exercise: %w", err)
	}

	return id.String(), nil
}

// Returns Exercises grouped by first letter
func (es *exerciseService) Fetch(c context.Context) ([]domain.Exercise, error) {
	ctx, cancel := context.WithTimeout(c, es.contextTimeout)
	defer cancel()

	res, err := es.exerciseRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}
