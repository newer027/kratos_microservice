package dao

import (
	"context"

	"github.com/go-kratos/kratos/pkg/cache"
	"github.com/go-kratos/kratos/pkg/log"
	"github.com/newer027/kratos_microservice/apps/reviews/internal/model"
)

// Review get data from cache if miss will call source method, then add to cache.
func (d *dao) Review(c context.Context, id int64) (reviews []*model.Review, err error) {
	var (
		_emptContest = make([]*model.Review, 0)
	)
	if reviews, err = d.GetCRecent(c, id); err != nil || len(reviews) == 0 {
		cache.MetricMisses.Inc("bts:Review")
		err = nil
		if reviews, err = d.RawReview(c, id); err != nil {
			log.Error("reviews.dao.Review error(%v)", err)
			return
		}
		if len(reviews) == 0 {
			reviews = _emptContest
			return
		}
		d.AddCRecent(c, id, reviews)
	}
	cache.MetricHits.Inc("bts:Review")
	return
}
