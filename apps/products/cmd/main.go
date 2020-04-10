package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/newer027/kratos_microservice/apps/products/internal/di"
	"github.com/go-kratos/kratos/pkg/naming/discovery"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden/resolver"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/net/trace/zipkin"
	"github.com/go-kratos/kratos/pkg/log"
)

func main() {
	flag.Parse()
	log.Init(nil) // debug flag: log.dir={path}
	defer log.Close()
	log.Info("products servive start")
	paladin.Init()

    zipkin.Init(&zipkin.Config{
        Endpoint: "http://127.0.0.1:9411/api/v2/spans",
	})
	resolver.Register(discovery.Builder())
	
	_, closeFunc, err := di.InitApp()
	if err != nil {
		panic(err)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			closeFunc()
			log.Info("products exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
