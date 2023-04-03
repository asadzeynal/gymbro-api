package service

import "github.com/asadzeynal/gymbro-api/internal/domain"

type exerciseService struct {
	exerciseRepo domain.ExerciseRepository
}

func NewExerciseService(er domain.ExerciseRepository) *exerciseService {
	return &exerciseService{er}
}

// Returns Exercises grouped by first letter
func FetchGroupedByInitial() {

}
