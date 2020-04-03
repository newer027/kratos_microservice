package dao

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/pkg/cache/memcache"
	"github.com/go-kratos/kratos/pkg/log"
	"github.com/newer027/kratos_microservice/apps/details/internal/model"
)

var _ _mc

// CacheDetail get data from mc
func (d *dao) CacheDetail(c context.Context, id int64) (res *model.Detail, err error) {
	key := keyDet(id)
	res = &model.Detail{}
	if err = d.mc.Get(c, key).Scan(res); err != nil {
		res = nil
		if err == memcache.ErrNotFound {
			err = nil
		}
	}
	if err != nil {
		log.Errorv(c, log.KV("CacheDetail", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// AddCacheDetail Set data to mc
func (d *dao) AddCacheDetail(c context.Context, id int64, val *model.Detail) (err error) {
	if val == nil {
		return
	}
	key := keyDet(id)
	item := &memcache.Item{Key: key, Object: val, Expiration: d.demoExpire, Flags: memcache.FlagJSON}
	if err = d.mc.Set(c, item); err != nil {
		log.Errorv(c, log.KV("AddCacheDetail", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// DeleteDetailCache delete data from mc
func (d *dao) DeleteDetailCache(c context.Context, id int64) (err error) {
	key := keyDet(id)
	if err = d.mc.Delete(c, key); err != nil {
		if err == memcache.ErrNotFound {
			err = nil
			return
		}
		log.Errorv(c, log.KV("DeleteDetailCache", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}
