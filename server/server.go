package server

import (
	"context"
	"github.com/gobuffalo/pop/v6"
	"github.com/kzh/arcane/pkg/arcanepb"
	"github.com/kzh/arcane/pkg/database"
	"github.com/kzh/arcane/pkg/models"

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

func (s *Server) KVGet(ctx context.Context, request *arcanepb.KVGetRequest) (*arcanepb.KVGetResponse, error) {
	kv, err := models.KVs.ByKey(s.conn, request.Key)
	if err != nil {
		return nil, err
	}
	return &arcanepb.KVGetResponse{
		Key:   kv.Key,
		Value: kv.Value,
	}, nil
}

func (s *Server) KVPut(ctx context.Context, request *arcanepb.KVPutRequest) (*arcanepb.KVPutResponse, error) {
	kv := &models.KV{
		Key:   request.Key,
		Value: request.Value,
	}
	err := kv.Create(s.conn)
	return &arcanepb.KVPutResponse{}, err
}
