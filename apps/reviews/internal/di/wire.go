// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"github.com/newer027/kratos_microservice/apps/reviews/internal/server/http"

	"github.com/newer027/kratos_microservice/apps/reviews/internal/service"

	"github.com/newer027/kratos_microservice/apps/reviews/internal/dao"

	"github.com/newer027/kratos_microservice/apps/reviews/internal/server/grpc"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
