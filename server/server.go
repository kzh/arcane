package server

import (
	"context"
	"github.com/gobuffalo/pop/v6"
	"github.com/kzh/arcane/pkg/database"

	"github.com/kzh/arcane/pkg/arcanepb"

	"go.uber.org/zap"
)

var _ arcanepb.ArcaneServer = (*Server)(nil)

type Server struct {
	conn *pop.Connection

	arcanepb.UnimplementedArcaneServer
}

func NewServer() (*Server, error) {
	conn, err := database.Connect()
	if err != nil {
		return nil, err
	}

	if err := database.Migrate(conn); err != nil {
		return nil, err
	}

	return &Server{
		conn: conn,
	}, nil
}

func (s *Server) Status(ctx context.Context, request *arcanepb.StatusRequest) (*arcanepb.StatusResponse, error) {
	zap.L().Info("Status", zap.String("foo", "bar"))

	return &arcanepb.StatusResponse{
		Status: arcanepb.StatusResponse_HEALTHY,
	}, nil
}

func (s *Server) KVGet(context.Context, *arcanepb.KVGetRequest) (*arcanepb.KVGetResponse, error) {
	return &arcanepb.KVGetResponse{}, nil
}

func (s *Server) KVPut(context.Context, *arcanepb.KVPutRequest) (*arcanepb.KVPutResponse, error) {
	return &arcanepb.KVPutResponse{}, nil
}
