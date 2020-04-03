
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

* **访问接口**： http://localhost:8000/product/id?id=1
* **discovery**: http://localhost:7171/
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
   - DISCOVERY_NODES=127.0.0.1:7171
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