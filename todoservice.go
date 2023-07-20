package todo

import (
	"context"
	"log"

	todoservice "github.com/TsubasaKanemitsu/golang-todo-app/gen/todoservice"
	"github.com/TsubasaKanemitsu/golang-todo-app/injector"
	"github.com/TsubasaKanemitsu/golang-todo-app/pkg/config"
	"github.com/TsubasaKanemitsu/golang-todo-app/usecase"
	"github.com/TsubasaKanemitsu/golang-todo-app/usecase/input"
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
	m, err := input.ToMTodoTaskInfo(p)
	if err != nil {
		s.logger.Fatalf("error: input.ToMTodoTaskInfo()")
		return false, err
	}
	return s.usecase.AddTodoTask(ctx, m)
}

// 指定したTodoタスクの詳細を取得する。
func (s *todoservicesrvc) GetTodoTask(ctx context.Context, p *todoservice.GetTodoTaskPayload) (res *todoservice.TodoTaskInfo, err error) {
	s.logger.Print("todoservice.GetTodoTask")

	m, err := s.usecase.GetTodoTask(ctx, p.ID)
	if err != nil {
		s.logger.Fatalf("error: usecase.GetTodoTask()")
		return &todoservice.TodoTaskInfo{}, err
	}
	startDate := m.StartDate.Format("2006-01-02")
	endDate := m.EndDate.Format("2006-01-02")

	res = &todoservice.TodoTaskInfo{
		ID:        m.ID,
		Title:     m.Title,
		Contents:  &m.Contents.String,
		Status:    m.Status,
		Label:     &m.Label.String,
		Asignee:   &m.Assignee.String,
		StartDate: &startDate,
		EndDate:   &endDate,
	}
	return res, nil
}

// Todoタスク一覧を取得する。
func (s *todoservicesrvc) GetTodoTaskList(ctx context.Context) (res []*todoservice.TodoTaskTitle, err error) {
	s.logger.Print("todoservice.GetTodoTaskList")
	mList, err := s.usecase.GetTodoTaskList(ctx)
	if err != nil {
		s.logger.Fatalf("error: usecase.GetTodoTaskList()")
		return []*todoservice.TodoTaskTitle{}, err
	}
	var response []*todoservice.TodoTaskTitle
	for _, m := range mList {
		t := &todoservice.TodoTaskTitle{
			ID:     m.ID,
			Title:  m.Title,
			Status: m.Status,
			Label:  &m.Label.String,
		}
		response = append(response, t)
	}
	return response, nil
}

// 指定したTodoタスクを更新する。
func (s *todoservicesrvc) UpdateTodoTask(ctx context.Context, p *todoservice.UpdateTodoTaskPayload) (res bool, err error) {
	s.logger.Print("todoservice.UpdateTodoTask")
	return
}

// 指定したTodoタスクを削除する。
func (s *todoservicesrvc) DeleteTodoTask(ctx context.Context, p *todoservice.DeleteTodoTaskPayload) (res bool, err error) {
	s.logger.Print("todoservice.DeleteTodoTask")
	return s.usecase.DeleteTodoTask(ctx, p.ID)
}
