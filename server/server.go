package server

import (
	"context"

	"github.com/kzh/arcane/pkg/arcanepb"

	"go.uber.org/zap"
)

var _ arcanepb.ArcaneServer = (*Server)(nil)

type Server struct {
	arcanepb.UnimplementedArcaneServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Status(ctx context.Context, request *arcanepb.StatusRequest) (*arcanepb.StatusResponse, error) {
	zap.L().Info("Status", zap.String("foo", "bar"))

	return &arcanepb.StatusResponse{
		Status: arcanepb.StatusResponse_HEALTHY,
	}, nil
}
