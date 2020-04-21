
## 准备

安装docker,go,[jsonnet](https://jsonnet.org/)

## 快速开始
下载项目
```bash
    git clone https://github.com/newer027/kratos_microservice.git
    cd kratos_microservice
    git submodule init
    git submodule update
    make docker-compose
```

* **访问接口**： http://localhost:8080/products/id?id=1
* **discovery**: http://localhost:7171/discovery/fetch/all
* **grafana**: http://localhost:3000/ 
* **jaeger**: http://localhost:16686/search
* **Prometheus**: http://localhost:9090/graph
* **AlertManager**: http://localhost:9093




/Users/Jacob/Desktop/创新小组个人事务/1_kratos_microservice/apps

env:
  global:
   - GO111MODULE=on
   - REGION=sh
   - ZONE=sh001
   - DEPLOY_ENV=dev
   - DISCOVERY_NODES=discovery:7171
   - HTTP_PERF=tcp://0.0.0.0:0
   - DOCKER_COMPOSE_VERSION=1.24.1
   - ZK_VERSION=3.5.6

go build

./cmd -conf ../configs --discovery.nodes="127.0.0.1:7171"
./cmd -conf ../configs -grpc=tcp://0.0.0.0:9000/?timeout=1s&idle_timeout=60s

lsof -i tcp:8888
lsof -i tcp:8000

docker-compose -f deployments/docker-mysql-memcache.yml up --build -d
docker-compose -f deployments/docker-mysql-memcache.yml down


调试details:
make docker-compose-dis
cd apps/details/details
go build;
./cmd -conf ../configs -grpc=tcp://0.0.0.0:9000/?timeout=1s&idle_timeout=60s -log.dir=/details_log
ps aux|grep cmd
docker-compose exec details ls /app/bin


调试products:
./cmd -conf ../configs --discovery.nodes="discovery:7171"
~/go/bin/grpcui -plaintext 127.0.0.1:9000


调试prometheus:
https://github.com/scotwells/prometheus-by-example

scrape_configs:
  - job_name: consul
    consul_sd_configs:
      - server: consul:8500
        datacenter: dc1
        tags:
          - http
    relabel_configs:
      - source_labels: [__meta_consul_service]
        target_label: app
  - job_name: grafana
    static_configs:
      - targets:
          - "grafana:3000"

调试zipkin tracing:
trace.SetGlobalTracer(trace.NewTracer(env.AppID, newReport(c), c.DisableSample))


0415 调试
{"code":0,"message":"0","ttl":1,"data":{"id":3,"name":"","price":0.5,"created_time":"2019-07-31T19:44:08Z"}}   没有name


topic = "AccountLog-T"

wake on lan

go clean --modcache

```
// UpdateExp update user exp.
func (s *Service) UpdateExp(c context.Context, arg *model.ArgAddExp) (err error) {
	var base *model.BaseInfo
	if base, err = s.BaseInfo(c, arg.Mid); err != nil {
		log.Error("s.BaseInfo(%d) error(%v)", arg.Mid, err)
		return
	}
	if base.Rank < 10000 {
		err = ecode.UserNoMember
		return
	}
	if arg.Count == 0 {
		log.Info("s.UpdateExp(%d) arg(%+v) count eq(0) continue", arg.Mid, arg)
		return
	}
	var exp int64
	if exp, err = s.mbDao.Exp(c, arg.Mid); err != nil {
		log.Error("s.mbDao.Exp(%d) error(%v)", arg.Mid, err)
		return
	}
	if exp == 0 {
		if _, err = s.mbDao.SetExp(c, arg.Mid, int64(arg.Count*model.ExpMulti)); err != nil {
			log.Error("s.mbDao.SetExp(%d) error(%v)", arg.Mid, err)
			return
		}
	} else {
		if _, err = s.mbDao.UpdateExp(c, arg.Mid, int64(arg.Count*model.ExpMulti)); err != nil {
			log.Error("s.mbDao.UpdateExp(%d) error(%v)", arg.Mid, err)
			return
		}
	}
	if err = s.mbDao.AddExplog(c, arg.Mid, exp/model.ExpMulti, (int64(arg.Count*model.ExpMulti)+exp)/model.ExpMulti, arg.Operate, arg.Reason, arg.IP); err != nil {
		log.Error("s.mbDao.AddExplog(%d) fromExp(%d) toExp(%d) oper(%s) reason(%s) ip(%s) error(%v)", arg.Mid, exp/model.ExpMulti, (int64(arg.Count*model.ExpMulti)+exp)/model.ExpMulti, arg.Operate, arg.Reason, arg.IP, err)
	} else {
		log.Info("s.mbDao.AddExplog(%d) fromExp(%d) toExp(%d) oper(%s) reason(%s) ip(%s)", arg.Mid, exp/model.ExpMulti, (int64(arg.Count*model.ExpMulti)+exp)/model.ExpMulti, arg.Operate, arg.Reason, arg.IP)
	}
	return
}
```