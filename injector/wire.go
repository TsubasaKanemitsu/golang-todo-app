//go:build wireinject
// +build wireinject

package injector

import (
	infra "github.com/TsubasaKanemitsu/golang-todo-app/infra/repository"
	"github.com/TsubasaKanemitsu/golang-todo-app/pkg/config"
	"github.com/TsubasaKanemitsu/golang-todo-app/pkg/db"
	"github.com/TsubasaKanemitsu/golang-todo-app/usecase"
	"github.com/google/wire"
)

func InitializeTodoService(c config.Postgres) (_ usecase.Todo) {
	wire.Build(
		db.NewPSQL,
		infra.NewTodoRepository,
		usecase.NewTodoUsecase,
	)
	return
}
