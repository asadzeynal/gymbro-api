package domain

import (
	"context"

	"github.com/gofrs/uuid/v5"
)

type TemplateSet struct {
	Id                 uuid.UUID
	TemplateExerciseId string
	Reps               int32
	Weight             float32
}

type TemplateExericse struct {
	Id           uuid.UUID
	TemplateId   uuid.UUID
	ExerciseId   uuid.UUID
	DisplayOrder int32
	Sets         []TemplateSet
}

type Template struct {
	Id        uuid.UUID
	Name      string
	UserId    uuid.UUID
	Exercises []TemplateExericse
}

type TemplateRepository interface {
	Store(ctx context.Context, template Template) (uuid.UUID, error)
	FetchByUserId(ctx context.Context, id uuid.UUID) (Exercise, error)
}
