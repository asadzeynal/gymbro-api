package repository

import (
	db "github.com/asadzeynal/gymbro-api/gen/sqlc"
	postgres "github.com/asadzeynal/gymbro-api/gen/sqlc"
	"github.com/asadzeynal/gymbro-api/internal/domain"
)

func templateSetToDomain(s db.TemplateSet) domain.TemplateSet {
	return domain.TemplateSet{
		Id:                 s.ID,
		TemplateExerciseId: s.TemplateExerciseID,
		Reps:               s.Reps,
		Weight:             float32(s.Weight),
	}
}

func templateExerciseToDomain(e db.TemplateExercise) domain.TemplateExericse {
	return domain.TemplateExericse{
		Id:           e.ID,
		TemplateId:   e.TemplateID,
		ExerciseId:   e.ExerciseID,
		DisplayOrder: e.DisplayOrder,
	}
}

func templateToDomain(t db.Template) domain.Template {
	return domain.Template{
		Id:     t.ID,
		Name:   t.Name,
		UserId: t.UserID,
	}
}

func exerciseToDomain(e postgres.Exercise) domain.Exercise {
	return domain.Exercise{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description,
	}
}
