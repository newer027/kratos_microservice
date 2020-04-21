package dao

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/pkg/cache/memcache"
	"github.com/go-kratos/kratos/pkg/log"
	"github.com/newer027/kratos_microservice/apps/ratings/internal/model"
)

var _ _mc

// CacheRating get data from mc
func (d *dao) CacheRating(c context.Context, id int64) (res *model.Rating, err error) {
	key := keyRat(id)
	res = &model.Rating{}
	if err = d.mc.Get(c, key).Scan(res); err != nil {
		res = nil
		if err == memcache.ErrNotFound {
			err = nil
		}
	}
	if err != nil {
		log.Errorv(c, log.KV("CacheRating", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// AddCacheRating Set data to mc
func (d *dao) AddCacheRating(c context.Context, id int64, val *model.Rating) (err error) {
	if val == nil {
		return
	}
	key := keyRat(id)
	item := &memcache.Item{Key: key, Object: val, Expiration: d.demoExpire, Flags: memcache.FlagJSON}
	if err = d.mc.Set(c, item); err != nil {
		log.Errorv(c, log.KV("AddCacheRating", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// DeleteRatingCache delete data from mc
func (d *dao) DeleteRatingCache(c context.Context, id int64) (err error) {
	key := keyRat(id)
	if err = d.mc.Delete(c, key); err != nil {
		if err == memcache.ErrNotFound {
			err = nil
			return
		}
		log.Errorv(c, log.KV("DeleteRatingCache", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}
