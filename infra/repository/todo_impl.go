package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/TsubasaKanemitsu/golang-todo-app/domain/model"
	"github.com/TsubasaKanemitsu/golang-todo-app/domain/repository"
	"github.com/TsubasaKanemitsu/golang-todo-app/infra/transfer"
	"github.com/TsubasaKanemitsu/golang-todo-app/pkg/dbmodels"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type todo struct {
	DB *sql.DB
}

func NewTodoRepository(db *sql.DB) repository.Todo {
	return &todo{
		DB: db,
	}
}

func (t *todo) Add(ctx context.Context, m *model.TodoTaskInfo) (bool, error) {
	e := transfer.ToTodoTaskInfoEntity(m)

	err := e.Insert(ctx, t.DB, boil.Infer())
	if err != nil {
		return false, fmt.Errorf("failed to insert todo task")
	}
	return true, nil
}

func (t *todo) GetTodoTask(ctx context.Context, id int) (*model.TodoTaskInfo, error) {
	e, err := dbmodels.FindMTodoTask(ctx, t.DB, id)
	if err != nil {
		return nil, fmt.Errorf("failet to fetch todo task (id: %v)", id)
	}

	m := transfer.ToTodoTaskInfoDomain(e)
	return m, nil
}

func (t *todo) GetTodoTaskList(ctx context.Context) ([]*model.TodoTaskTitle, error) {
	entities, err := dbmodels.MTodoTasks().All(ctx, t.DB)
	if err != nil {
		return nil, fmt.Errorf("failet to fetch todo task list")
	}
	models := make([]*model.TodoTaskTitle, len(entities))
	for _, e := range entities {

		m := &model.TodoTaskTitle{
			ID:     e.ID,
			Title:  e.Title,
			Label:  e.Label,
			Status: e.Status,
		}
		models = append(models, m)
	}

	return models, nil
}

func (t *todo) DeleteTodoTask(ctx context.Context, id int) (bool, error) {
	e, err := dbmodels.FindMTodoTask(ctx, t.DB, id)
	if err != nil {
		return false, fmt.Errorf("failet to fetch todo task (id: %v)", id)
	}
	_, err = e.Delete(ctx, t.DB)
	if err != nil {
		return false, fmt.Errorf("failet to delete todo task")
	}
	log.Printf("success to delete todo task (id: %v)", e.ID)

	return true, nil
}
