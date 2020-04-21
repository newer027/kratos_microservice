package dao

import (
	"context"

	"github.com/newer027/kratos_microservice/apps/ratings/internal/model"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/database/sql"
	"github.com/go-kratos/kratos/pkg/log"
)

const (
	_queryRating = "select id,product_id,score,updated_time from ratings where id=?"
)

func NewDB() (db *sql.DB, cf func(), err error) {
	var (
		cfg sql.Config
		ct paladin.TOML
	)
	if err = paladin.Get("db.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Client").UnmarshalTOML(&cfg); err != nil {
		return
	}
	db = sql.NewMySQL(&cfg)
	cf = func() {db.Close()}
	return
}

func (d *dao) RawRating(ctx context.Context, id int64) (rating *model.Rating, err error) {
	rating = new(model.Rating)
	raw := d.db.QueryRow(ctx, _queryRating, id)
	if err = raw.Scan(&rating.ID, &rating.ProductID, &rating.Score, &rating.UpdatedTime); err != nil {
		log.Error("query rating err(%v)", err)
	}
	return
}

