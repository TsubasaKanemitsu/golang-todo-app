package repository

import (
	"context"

	"github.com/TsubasaKanemitsu/golang-todo-app/domain/model"
)

type Todo interface {
	Add(ctx context.Context, m *model.TodoTaskInfo) (ok bool, err error)
	GetTodoTask(ctx context.Context, id int) (m *model.TodoTaskInfo, err error)
	GetTodoTaskList(ctx context.Context) (mList []*model.TodoTaskTitle, err error)
	//UpdateTodoTask(ctx context.Context, )
	DeleteTodoTask(ctx context.Context, id int) (ok bool, err error)
}
