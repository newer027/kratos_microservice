package http

import (
	"net/http"
	"strconv"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/newer027/kratos_microservice/apps/products/internal/service"

	"github.com/go-kratos/kratos/pkg/log"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
)

var svc *service.Service

// New new a bm server.
func New(s *service.Service) (engine *bm.Engine, err error) {
	var (
		cfg bm.ServerConfig
		ct  paladin.TOML
	)
	if err = paladin.Get("http.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Server").UnmarshalTOML(&cfg); err != nil {
		return
	}
	svc = s
	engine = bm.DefaultServer(&cfg)
	initRouter(engine)
	err = engine.Start()
	return
}

func initRouter(e *bm.Engine) {
	e.Ping(ping)
	g := e.Group("/products")
	{
		g.GET("/id", get)
	}
}

func ping(ctx *bm.Context) {
	if _, err := svc.Ping(ctx, nil); err != nil {
		log.Error("ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

// example for http request handler.
func get(ctx *bm.Context) {
	params := ctx.Request.Form
	ID, err := strconv.ParseInt(params.Get("id"), 10, 64)
	if err != nil {
		log.Warn("%v", err)
		return
	}
	p, err := svc.Get(ctx, ID)
	if err != nil {
		log.Error("get product by id error(%v)", err)
		return
	}
	ctx.JSON(p, nil)
}
