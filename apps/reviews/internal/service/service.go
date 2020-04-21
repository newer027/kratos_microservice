package service

import (
	"context"
	"github.com/pkg/errors"
	"github.com/golang/protobuf/ptypes"

	pb "github.com/newer027/kratos_microservice/apps/reviews/api"
	"github.com/newer027/kratos_microservice/apps/reviews/internal/dao"
	"github.com/newer027/kratos_microservice/apps/reviews/internal/model"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/wire"
)

var Provider = wire.NewSet(New, wire.Bind(new(pb.ReviewsServer), new(*Service)))

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
func (s *Service) Query(ctx context.Context, req *pb.QueryReviewsRequest) (resp *pb.QueryReviewsResponse, err error) {
	resp = &pb.QueryReviewsResponse{
		Reviews: make([]*pb.Review, 0),
	}
	var res []*model.Review
	if res, err = s.dao.Review(ctx, req.ProductID); err != nil {
		return nil, errors.Wrap(err, "Reviews service get reviews error")
	}

	for _, v := range res {
		ut, err := ptypes.TimestampProto(v.CreatedTime)
		if err != nil {
			return nil, err
		}
		tmp := &pb.Review{
			Id:          	int64(v.ID),
			ProductID:      v.ProductID,
			Message:       	v.Message,
			CreatedTime: 	ut,
		}
		resp.Reviews = append(resp.Reviews, tmp)
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
