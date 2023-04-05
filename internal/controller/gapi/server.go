package gapi

import "github.com/asadzeynal/gymbro-api/gen/pb"

type Server struct {
	pb.UnimplementedGymBroServer
}

func NewServer() *Server {
	return &Server{}
}
