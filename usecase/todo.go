package usecase

import (
	"context"

	"github.com/TsubasaKanemitsu/golang-todo-app/domain/model"
)

type Todo interface {
	AddTodoTask(ctx context.Context, m *model.TodoTaskInfo) (bool, error)
	GetTodoTask(ctx context.Context, id int) (m *model.TodoTaskInfo, err error)
	GetTodoTaskList(ctx context.Context) (mList []*model.TodoTaskTitle, err error)
	UpdateTodoTask(ctx context.Context, m *model.TodoTaskInfo) (bool, error)
	DeleteTodoTask(ctx context.Context, id int) (ok bool, err error)
}
