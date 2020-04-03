package service

import (
	"context"
	"github.com/newer027/kratos_microservice/apps/products/internal/dao"
	"github.com/newer027/kratos_microservice/apps/products/internal/model"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/go-kratos/kratos/pkg/log"
)

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

// Get Get
func (s *Service) Get(ctx context.Context, id int64) (detail *model.Detail, err error) {
	detail, err = s.dao.GetDetail(ctx, id)
	if err != nil {
		log.Warnv(ctx, log.KV("log", "get detail fail: err="+err.Error()))
		return
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
