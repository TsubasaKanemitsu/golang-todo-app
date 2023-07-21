package usecase

import (
	"context"
	"fmt"

	"github.com/TsubasaKanemitsu/golang-todo-app/domain/model"
	"github.com/TsubasaKanemitsu/golang-todo-app/domain/repository"
	"github.com/go-playground/validator/v10"
)

type todo struct {
	repository repository.Todo
}

func NewTodoUsecase(r repository.Todo) Todo {
	return &todo{repository: r}
}

func (u *todo) AddTodoTask(ctx context.Context, m *model.TodoTaskInfo) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, fmt.Errorf("validation error: %v", err)
	}
	return u.repository.Add(ctx, m)

}

func (u *todo) GetTodoTask(ctx context.Context, id int) (*model.TodoTaskInfo, error) {
	return u.repository.GetTodoTask(ctx, id)
}

func (u *todo) GetTodoTaskList(ctx context.Context) ([]*model.TodoTaskTitle, error) {
	return u.repository.GetTodoTaskList(ctx)
}

func (u *todo) UpdateTodoTask(ctx context.Context, m *model.TodoTaskInfo) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, fmt.Errorf("validation error: %v", err)
	}
	return u.repository.UpdateTodoTask(ctx, m)

}

func (u *todo) DeleteTodoTask(ctx context.Context, id int) (bool, error) {
	return u.repository.DeleteTodoTask(ctx, id)
}
