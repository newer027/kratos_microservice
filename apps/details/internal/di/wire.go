// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"github.com/newer027/kratos_microservice/apps/details/internal/server/http"

	"github.com/newer027/kratos_microservice/apps/details/internal/service"

	"github.com/newer027/kratos_microservice/apps/details/internal/dao"

	"github.com/newer027/kratos_microservice/apps/details/internal/server/grpc"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
