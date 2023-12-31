// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injector

import (
	"github.com/TsubasaKanemitsu/golang-todo-app/infra/repository"
	"github.com/TsubasaKanemitsu/golang-todo-app/pkg/config"
	"github.com/TsubasaKanemitsu/golang-todo-app/pkg/db"
	"github.com/TsubasaKanemitsu/golang-todo-app/usecase"
)

// Injectors from wire.go:

func InitializeTodoService(c config.Postgres) usecase.Todo {
	sqlDB := db.NewPSQL(c)
	todo := repository.NewTodoRepository(sqlDB)
	usecaseTodo := usecase.NewTodoUsecase(todo)
	return usecaseTodo
}
