package todo

import (
	"context"
	"log"

	"github.com/TsubasaKanemitsu/golang-todo-app/domain/model"
	todoservice "github.com/TsubasaKanemitsu/golang-todo-app/gen/todoservice"
	"github.com/TsubasaKanemitsu/golang-todo-app/injector"
	"github.com/TsubasaKanemitsu/golang-todo-app/pkg/config"
	"github.com/TsubasaKanemitsu/golang-todo-app/usecase"
)

// todoservice service example implementation.
// The example methods log the requests and return zero values.
type todoservicesrvc struct {
	logger  *log.Logger
	usecase usecase.Todo
}

// NewTodoservice returns the todoservice service implementation.
func NewTodoservice(logger *log.Logger, c config.Postgres) todoservice.Service {
	uc := injector.InitializeTodoService(c)
	return &todoservicesrvc{logger, uc}
}

// Todoタスクを追加する。
func (s *todoservicesrvc) AddTodoTask(ctx context.Context, p *todoservice.AddTodoTaskPayload) (res bool, err error) {
	s.logger.Print("todoservice.addTodoTask")
	s.usecase.AddTodoTask(ctx, m*model.TodoTaskInfo)
	return
}

// 指定したTodoタスクの詳細を取得する。
func (s *todoservicesrvc) GetTodoTask(ctx context.Context, p *todoservice.GetTodoTaskPayload) (res *todoservice.TodoTaskInfo, err error) {
	res = &todoservice.TodoTaskInfo{}
	s.logger.Print("todoservice.GetTodoTask")
	return
}

// Todoタスク一覧を取得する。
func (s *todoservicesrvc) GetTodoTaskList(ctx context.Context) (res []*todoservice.TodoTaskTitleList, err error) {
	s.logger.Print("todoservice.GetTodoTaskList")
	return
}

// 指定したTodoタスクを更新する。
func (s *todoservicesrvc) UpdateTodoTask(ctx context.Context, p *todoservice.UpdateTodoTaskPayload) (res bool, err error) {
	s.logger.Print("todoservice.UpdateTodoTask")
	return
}

// 指定したTodoタスクを削除する。
func (s *todoservicesrvc) DeleteTodoTask(ctx context.Context, p *todoservice.DeleteTodoTaskPayload) (res bool, err error) {
	s.logger.Print("todoservice.DeleteTodoTask")
	return
}
