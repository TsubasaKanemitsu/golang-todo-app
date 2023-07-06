package todo

import (
	"context"
	"log"

	service "github.com/TsubasaKanemitsu/golang-todo-app/gen/service"
)

// service service example implementation.
// The example methods log the requests and return zero values.
type servicesrvc struct {
	logger *log.Logger
}

// NewService returns the service service implementation.
func NewService(logger *log.Logger) service.Service {
	return &servicesrvc{logger}
}

// Todoタスクを追加する。
func (s *servicesrvc) AddTodoTask(ctx context.Context, p *service.AddTodoTaskPayload) (res bool, err error) {
	s.logger.Print("service.addTodoTask")
	return true, nil
}

// 指定したTodoタスクの詳細を取得する。
func (s *servicesrvc) GetTodoTask(ctx context.Context, p *service.GetTodoTaskPayload) (res *service.TodoTaskInfo, err error) {
	res = &service.TodoTaskInfo{}
	s.logger.Print("service.GetTodoTask")
	return
}

// Todoタスク一覧を取得する。
func (s *servicesrvc) GetTodoTaskList(ctx context.Context) (res []*service.TodoTaskTitleList, err error) {
	s.logger.Print("service.GetTodoTaskList")
	return
}

// 指定したTodoタスクを更新する。
func (s *servicesrvc) UpdateTodoTask(ctx context.Context, p *service.UpdateTodoTaskPayload) (res bool, err error) {
	s.logger.Print("service.UpdateTodoTask")
	return true, nil
}

// 指定したTodoタスクを削除する。
func (s *servicesrvc) DELETETodoTask(ctx context.Context, p *service.DELETETodoTaskPayload) (res bool, err error) {
	s.logger.Print("service.DELETETodoTask")
	return true, nil
}
