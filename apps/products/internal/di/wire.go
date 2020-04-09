// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"github.com/newer027/kratos_microservice/apps/products/internal/server/http"

	"github.com/newer027/kratos_microservice/apps/products/internal/service"

	"github.com/newer027/kratos_microservice/apps/products/internal/dao"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.New, http.New, NewApp))
}
