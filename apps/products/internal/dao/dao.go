package dao

import (
	"context"
	"time"
	"github.com/pkg/errors"

	"github.com/newer027/kratos_microservice/apps/products/internal/model"
	"github.com/go-kratos/kratos/pkg/cache/memcache"
	"github.com/go-kratos/kratos/pkg/cache/redis"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/database/sql"
	"github.com/go-kratos/kratos/pkg/sync/pipeline/fanout"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	xtime "github.com/go-kratos/kratos/pkg/time"
	details "github.com/newer027/kratos_microservice/apps/details/api"
	"github.com/golang/protobuf/ptypes"

	"github.com/google/wire"
)

var Provider = wire.NewSet(New, NewDB, NewRedis, NewMC)

//go:generate kratos tool genbts
// Dao dao interface
type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)
	// bts: -nullcache=&model.Article{ID:-1} -check_null_code=$!=nil&&$.ID==-1
	GetDetail(c context.Context, id int64) (*model.Detail, error)
}

// dao dao.
type dao struct {
	db          *sql.DB
	redis       *redis.Redis
	mc          *memcache.Memcache
	cache *fanout.Fanout
	demoExpire int32
	detailsClient     details.DetailsClient
}

//GRPCConf .
type GRPCConf struct {
	WardenConf *warden.ClientConfig
	Addr       string
}

// New new a dao and return.
func New(r *redis.Redis, mc *memcache.Memcache, db *sql.DB) (d Dao, cf func(), err error) {
	return newDao(r, mc, db)
}

func newDao(r *redis.Redis, mc *memcache.Memcache, db *sql.DB) (d *dao, cf func(), err error) {
	
	var cfg struct{
		DemoExpire xtime.Duration
		GRPCClient map[string]*GRPCConf
	}
	if err = paladin.Get("application.toml").UnmarshalTOML(&cfg); err != nil {
		return
	}
	d = &dao{
		db: db,
		redis: r,
		mc: mc,
		cache: fanout.New("cache"),
		demoExpire: int32(time.Duration(cfg.DemoExpire) / time.Second),
		detailsClient:     newDetailsClient(cfg.GRPCClient["details"]),
	}
	cf = d.Close
	return
}

// Close close the resource.
func (d *dao) Close() {
	d.cache.Close()
}

// Ping ping the resource.
func (d *dao) Ping(ctx context.Context) (err error) {
	return nil
}

// newNoticeClient .
func newDetailsClient(cfg *GRPCConf) details.DetailsClient {
	cc, err := warden.NewClient(cfg.WardenConf).Dial(context.Background(), cfg.Addr)
	if err != nil {
		panic(err)
	}
	return details.NewDetailsClient(cc)
}

// GetDetail GetDetail
func (d *dao) GetDetail(ctx context.Context, id int64) (detail *model.Detail, err error) {
	req := &details.GetDetailRequest{Id: id}
	res, err := d.detailsClient.Get(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "get detail error")
	}
	ct, err := ptypes.Timestamp(res.CreatedTime)

	detail = &model.Detail{
		ID:          res.Id,
		Price:       res.Price,
		CreatedTime: ct,
	}
	return
}