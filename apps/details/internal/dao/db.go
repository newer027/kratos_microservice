package dao

import (
	"context"

	"github.com/newer027/kratos_microservice/apps/details/internal/model"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/database/sql"
	"github.com/go-kratos/kratos/pkg/log"
)

const (
	_queryDetail = "select `id`,`name`,`price`,`created_time` from check_task where `id` = ?"
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

func (d *dao) RawDetail(ctx context.Context, id int64) (det *model.Detail, err error) {
	raw := d.db.QueryRow(ctx, _queryDetail, id)
	if err = raw.Scan(&det.ID, &det.Name, &det.Price, &det.CreatedTime); err != nil {
		log.Error("query detail err(%v)", err)
	}
	return
}
