package gapi

import (
	"context"

	pb "github.com/asadzeynal/gymbro-api/gen/proto/exercise/v1"
	"github.com/asadzeynal/gymbro-api/internal/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type exerciseService interface {
	Store(c context.Context, ex *domain.Exercise) (string, error)
	Fetch(c context.Context) ([]domain.Exercise, error)
}

type Server struct {
	pb.UnimplementedExerciseServiceServer
	es exerciseService
}

func NewServer(es exerciseService) *Server {
	return &Server{es: es}
}

func (s *Server) CreateExercise(c context.Context, req *pb.CreateExerciseRequest) (*pb.CreateExerciseResponse, error) {
	exercise := domain.Exercise{
		Description: req.Description,
		Name:        req.Name,
	}
	id, err := s.es.Store(c, &exercise)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error while creating the exercise: %w", err)
	}
	return &pb.CreateExerciseResponse{Id: id}, nil
}

func (s *Server) GetExercises(c context.Context, req *pb.GetExercisesRequest) (*pb.GetExercisesResponse, error) {
	excs, err := s.es.Fetch(c)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error while fetching exercises: %w", err)
	}

	res := make([]*pb.Exercise, len(excs))
	for i := range excs {
		exercise := excs[i]
		res[i] = &pb.Exercise{
			Id:          exercise.ID.String(),
			Name:        exercise.Name,
			Description: exercise.Description,
		}
	}

	return &pb.GetExercisesResponse{Exercises: res}, nil
}
