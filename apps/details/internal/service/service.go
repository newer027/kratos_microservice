package service

import (
	"context"
	"github.com/pkg/errors"
	"github.com/golang/protobuf/ptypes"

	pb "github.com/newer027/kratos_microservice/apps/details/api"
	"github.com/newer027/kratos_microservice/apps/details/internal/dao"
	"github.com/newer027/kratos_microservice/apps/details/internal/model"

	"github.com/go-kratos/kratos/pkg/conf/paladin"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/wire"
)

var Provider = wire.NewSet(New, wire.Bind(new(pb.DetailsServer), new(*Service)))

// Service service.
type Service struct {
	ac  *paladin.Map
	dao dao.Dao
}

// New new a service and return.
func New(d dao.Dao) (s *Service, cf func(), err error) {
	s = &Service{
		ac:  &paladin.TOML{},
		dao: d,
	}
	cf = s.Close
	err = paladin.Watch("application.toml", s.ac)
	return
}

// Get bm demo func.
func (s *Service) Get(ctx context.Context, req *pb.GetDetailRequest) (resp *pb.Detail, err error) {
	var p *model.Detail
	if p, err = s.dao.Get(ctx, req.Id); err != nil {
		return nil, errors.Wrap(err, "details service get detail error")
	}
	ut, err := ptypes.TimestampProto(p.CreatedTime)
	resp = &pb.Detail{
		Id:          int64(p.ID),
		Name:        p.Name,
		Price:       p.Price,
		CreatedTime: ut,
	}
	return
}

// Close close the resource.
func (s *Service) Close() {
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, s.dao.Ping(ctx)
}