package domain

import (
	"context"

	"github.com/gofrs/uuid/v5"
)

type TemplateSet struct {
	Id                 string
	TemplateExerciseId string
	Reps               int32
	Weight             float32
}

type TemplateExericse struct {
	Id           string
	TemplateId   string
	ExerciseId   string
	DisplayOrder int32
	Sets         []TemplateSet
}

type Template struct {
	Id        string
	Name      string
	UserId    string
	Exercises []TemplateExericse
}

type TemplateRepository interface {
	Store(ctx context.Context, template Template) (uuid.UUID, error)
	FetchByUserId(ctx context.Context, id uuid.UUID) (Exercise, error)
}
