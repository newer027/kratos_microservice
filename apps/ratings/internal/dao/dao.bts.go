package dao

import (
	"context"

	"github.com/go-kratos/kratos/pkg/cache"
	"github.com/newer027/kratos_microservice/apps/ratings/internal/model"
)

// Article get data from cache if miss will call source method, then add to cache.
func (d *dao) Get(c context.Context, id int64) (res *model.Rating, err error) {
	addCache := true
	res, err = d.CacheRating(c, id)
	if err != nil {
		addCache = false
		err = nil
	}
	defer func() {
		if res != nil && res.ID == -1 {
			res = nil
		}
	}()
	if res != nil {
		cache.MetricHits.Inc("bts:Rating")
		return
	}
	cache.MetricMisses.Inc("bts:Rating")
	res, err = d.RawRating(c, id)
	if err != nil {
		return
	}
	miss := res
	if miss == nil {
		miss = &model.Rating{ID: -1}
	}
	if !addCache {
		return
	}
	d.cache.Do(c, func(c context.Context) {
		d.AddCacheRating(c, id, miss)
	})
	return
}
