package dao

import (
	"fmt"
	"context"

	"github.com/newer027/kratos_microservice/apps/reviews/internal/model"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/database/sql"
	"github.com/go-kratos/kratos/pkg/log"
)

const (
	_contestsSQL     = "SELECT id,product_id,message,created_time FROM reviews WHERE product_id=%d ORDER BY ID ASC"
)

//NewDB NewDB
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

func (d *dao) RawReview(ctx context.Context, id int64) (reviews []*model.Review, err error) {
	var (
		rows *sql.Rows
	)
	var (
		_emptContest = make([]*model.Review, 0)
	)
	if rows, err = d.db.Query(ctx, fmt.Sprintf(_contestsSQL, id)); err != nil {
		log.Error("Contests: db.Exec(%d) error(%v)", id, err)
		reviews = _emptContest
		return
	}
	defer rows.Close()
	for rows.Next() {
		r := new(model.Review)
		if err = rows.Scan(&r.ID, &r.ProductID, &r.Message, &r.CreatedTime); err != nil {
			log.Error("Contests:row.Scan() error(%v)", err)
			return
		}
		reviews = append(reviews, r)
	}
	if err = rows.Err(); err != nil {
		log.Error("rows.Err() error(%v)", err)
	}
	return
}
