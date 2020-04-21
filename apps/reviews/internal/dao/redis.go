package dao

import (
	"fmt"
	"context"
	"encoding/json"

	"github.com/newer027/kratos_microservice/apps/reviews/internal/model"
	"github.com/go-kratos/kratos/pkg/cache/redis"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/log"
)

//NewRedis NewRedis
func NewRedis() (r *redis.Redis, cf func(), err error) {
	var (
		cfg redis.Config
		ct  paladin.Map
	)
	if err = paladin.Get("redis.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Client").UnmarshalTOML(&cfg); err != nil {
		return
	}
	r = redis.NewRedis(&cfg)
	cf = func() { r.Close() }
	return
}

func (d *dao) PingRedis(ctx context.Context) (err error) {
	if _, err = d.redis.Do(ctx, "SET", "ping", "pong"); err != nil {
		log.Error("conn.Set(PING) error(%v)", err)
	}
	return
}

// GetCRecent get contest recent data
func (d *dao) GetCRecent(ctx context.Context, id int64) (data []*model.Review, err error) {
	var (
		bs   []byte
		vals interface{}
	)
	key := fmt.Sprintf("review_%d", id)
	vals, err = d.redis.Do(ctx, "GET", key)
	if bs, err = redis.Bytes(vals, err); err != nil {
		data = nil
		if err == redis.ErrNil {
			err = nil
		} else {
			log.Error("GetCRecent conn.Do(GET,%d) error(%v)", id, err)
		}
		return
	}
	data = make([]*model.Review, 0)
	if err = json.Unmarshal(bs, &data); err != nil {
		log.Error("GetCRecent json.Unmarshal(%s) error(%v)", string(bs), err)
	}
	return
}

// AddCRecent add contest recent data
func (d *dao) AddCRecent(c context.Context, id int64, data []*model.Review) (err error) {
	var (
		bs   []byte
	)
	key := fmt.Sprintf("review_%d", id)
	if bs, err = json.Marshal(data); err != nil {
		log.Error("AddCRecent json.Marshal() error(%v)", err)
		return
	}
	if _, err = d.redis.Do(c, "SET", key, bs); err != nil {
		log.Error("AddCRecent conn.Send(SET,%d) error(%v)", id, err)
		return
	}
	if _, err = d.redis.Do(c, "EXPIRE", key, 60); err != nil {
		log.Error("AddCRecent conn.Send(EXPIRE,%d) error(%v)", id, err)
		return
	}
	return
}