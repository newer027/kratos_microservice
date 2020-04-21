package service

import (
	"context"
	"github.com/pkg/errors"
	"github.com/golang/protobuf/ptypes"

	pb "github.com/newer027/kratos_microservice/apps/ratings/api"
	"github.com/newer027/kratos_microservice/apps/ratings/internal/dao"
	"github.com/newer027/kratos_microservice/apps/ratings/internal/model"

	"github.com/go-kratos/kratos/pkg/conf/paladin"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/wire"
)

var Provider = wire.NewSet(New, wire.Bind(new(pb.RatingsServer), new(*Service)))

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
func (s *Service) Get(ctx context.Context, req *pb.GetRatingRequest) (resp *pb.Rating, err error) {
	var p *model.Rating
	if p, err = s.dao.Get(ctx, req.ProductID); err != nil {
		return nil, errors.Wrap(err, "Rating service get error")
	}
	ut, err := ptypes.TimestampProto(p.UpdatedTime)
	resp = &pb.Rating{
		Id:          int64(p.ID),
		ProductID:   p.ProductID,
		Score:       p.Score,
		UpdatedTime: ut,
	}
	return
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, s.dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
}
