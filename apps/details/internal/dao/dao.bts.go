package dao

import (
	"context"

	"github.com/go-kratos/kratos/pkg/cache"
	"github.com/newer027/kratos_microservice/apps/details/internal/model"
)

// Article get data from cache if miss will call source method, then add to cache.
func (d *dao) Get(c context.Context, id int64) (res *model.Detail, err error) {
	addCache := true
	res, err = d.CacheDetail(c, id)
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
		cache.MetricHits.Inc("bts:Detail")
		return
	}
	cache.MetricMisses.Inc("bts:Detail")
	res, err = d.RawDetail(c, id)
	if err != nil {
		return
	}
	miss := res
	if miss == nil {
		miss = &model.Detail{ID: -1}
	}
	if !addCache {
		return
	}
	d.cache.Do(c, func(c context.Context) {
		d.AddCacheDetail(c, id, miss)
	})
	return
}
