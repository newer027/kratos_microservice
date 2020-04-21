package dao

import (
	"context"
	"time"
	"github.com/pkg/errors"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	"github.com/go-kratos/kratos/pkg/sync/pipeline/fanout"
	xtime "github.com/go-kratos/kratos/pkg/time"
	"github.com/golang/protobuf/ptypes"
	details "github.com/newer027/kratos_microservice/apps/details/api"
	ratings "github.com/newer027/kratos_microservice/apps/ratings/api"
	reviews "github.com/newer027/kratos_microservice/apps/reviews/api"
	"github.com/newer027/kratos_microservice/apps/products/internal/model"
)

// Dao dao interface
type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)
	GetDetail(c context.Context, id int64) (*model.Detail, error)
	GetRating(c context.Context, id int64) (*model.Rating, error)
	GetReview(c context.Context, id int64) ([]*model.Review, error)
}

// dao dao.
type dao struct {
	cache         *fanout.Fanout
	demoExpire    int32
	detailsClient details.DetailsClient
	ratingsClient ratings.RatingsClient
	reviewsClient reviews.ReviewsClient

	
}

//GRPCConf .
type GRPCConf struct {
	WardenConf *warden.ClientConfig
	Addr       string
}

// New new a dao and return.
func New() (d Dao, cf func(), err error) {
	return newDao()
}

func newDao() (d *dao, cf func(), err error) {

	var cfg struct {
		DemoExpire xtime.Duration
		GRPCClient map[string]*GRPCConf
	}
	if err = paladin.Get("application.toml").UnmarshalTOML(&cfg); err != nil {
		return
	}
	d = &dao{
		cache:         fanout.New("cache"),
		demoExpire:    int32(time.Duration(cfg.DemoExpire) / time.Second),
		detailsClient: newDetailsClient(cfg.GRPCClient["details"]),
		ratingsClient: newRatingsClient(cfg.GRPCClient["ratings"]),
		reviewsClient: newReviewsClient(cfg.GRPCClient["reviews"]),
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

// newNoticeClient .
func newRatingsClient(cfg *GRPCConf) ratings.RatingsClient {
	cc, err := warden.NewClient(cfg.WardenConf).Dial(context.Background(), cfg.Addr)
	if err != nil {
		panic(err)
	}
	return ratings.NewRatingsClient(cc)
}

// newNoticeClient .
func newReviewsClient(cfg *GRPCConf) reviews.ReviewsClient {
	cc, err := warden.NewClient(cfg.WardenConf).Dial(context.Background(), cfg.Addr)
	if err != nil {
		panic(err)
	}
	return reviews.NewReviewsClient(cc)
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
		Name:        res.Name,
		Price:       res.Price,
		CreatedTime: ct,
	}
	return
} 

// GetDetail GetDetail
func (d *dao) GetRating(ctx context.Context, id int64) (rating *model.Rating, err error) {
	req := &ratings.GetRatingRequest{ProductID: id}
	res, err := d.ratingsClient.Get(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "get rating error")
	}
	ct, err := ptypes.Timestamp(res.UpdatedTime)

	rating = &model.Rating{
		ID:          res.Id,
		ProductID:   res.ProductID,
		Score:       res.Score,
		UpdatedTime: ct,
	}
	return
} 

// GetDetail GetDetail
func (d *dao) GetReview(ctx context.Context, id int64) (mreviews []*model.Review, err error) {
	req := &reviews.QueryReviewsRequest{ProductID: id}
	res, err := d.reviewsClient.Query(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "get reviews error")
	}
	for _, re := range res.Reviews {
		ct, err := ptypes.Timestamp(re.CreatedTime)
		if err != nil {
			return nil, errors.Wrap(err, "parse timestamp error")
		}
		review := &model.Review{
			ID:          re.Id,
			ProductID:   re.ProductID,
			Message:     re.Message,
			CreatedTime: ct,
		}
		mreviews = append(mreviews, review)
	}
	return
} 
/* 
const (
	_queryDetail = "select id,name,price,created_time from details where id=?"
)

func NewDB() (db *sql.DB, cf func(), err error) {
	var (
		cfg sql.Config
		ct  paladin.TOML
	)
	if err = paladin.Get("db.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Client").UnmarshalTOML(&cfg); err != nil {
		return
	}
	db = sql.NewMySQL(&cfg)
	cf = func() { db.Close() }
	return
}

func (d *dao) GetDetail(ctx context.Context, id int64) (det *model.Detail, err error) {
	det = new(model.Detail)
	err = d.db.QueryRow(ctx, _queryDetail, id).Scan(&det.ID, &det.Name, &det.Price, &det.CreatedTime)
	if err != nil {
		log.Error("query detail err(%v)", err)
	}
	return
}
 */