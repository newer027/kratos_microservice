package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-kratos/kratos/pkg/naming"
	"github.com/go-kratos/kratos/pkg/naming/discovery"
	"github.com/newer027/kratos_microservice/apps/details/internal/di"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/net/trace/zipkin"
	"github.com/go-kratos/kratos/pkg/log"
)

func main() {
	flag.Parse()
	log.Init(nil) // debug flag: log.dir={path}
	defer log.Close()
	log.Info("details service start")
	paladin.Init()
	_, closeFunc, err := di.InitApp()
	if err != nil {
		panic(err)
	}

    zipkin.Init(&zipkin.Config{
        Endpoint: "http://127.0.0.1:9411/api/v2/spans",
	})
	
	ip := "127.0.0.1"
	port := "9000"
	Zone := "sh001"
	DeployEnv := "dev"
	AppID := "Details"
	hn := "test"

	var cfg = &discovery.Config{
		Nodes:  []string{"127.0.0.1:7171"},
		Zone:   "sh001",
		Env:    "dev",
	}
	dis := discovery.New(cfg)

	ins := &naming.Instance{
		Zone:  Zone,
		Env:   DeployEnv,
		AppID: AppID,
		Hostname: hn,
		Addrs: []string{
			"grpc://" + ip + ":" + port,
		},
	}
	cancel, err := dis.Register(context.Background(), ins)
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
			if cancel != nil {
				cancel()
			}
			closeFunc()
			log.Info("details exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
