package grpc

import (
	pb "github.com/newer027/kratos_microservice/apps/details/api"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
)

// New new a grpc server.
func New(svc pb.DetailsServer) (ws *warden.Server, err error) {
	var (
		cfg warden.ServerConfig
		ct paladin.TOML
	)
	if err = paladin.Get("grpc.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Server").UnmarshalTOML(&cfg); err != nil {
		return
	}
	ws = warden.NewServer(&cfg)
	pb.RegisterDetailsServer(ws.Server(), svc)
	ws, err = ws.Start()
	return
}
