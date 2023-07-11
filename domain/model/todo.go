package model

import (
	"time"

	"github.com/TsubasaKanemitsu/golang-todo-app/pkg/dbmodels"
	"github.com/volatiletech/null/v8"
)

type TodoTaskInfo dbmodels.MTodoTask

func NewTodoTaskInfo(
	id int,
	title string,
	contents null.String,
	assignee null.String,
	label null.String,
	status string,
	startDate time.Time,
	endDate time.Time,
	createdAt null.Time,
	updatedAt null.Time,
) *TodoTaskInfo {
	return &TodoTaskInfo{
		ID:        id,
		Title:     title,
		Contents:  contents,
		Assignee:  assignee,
		Label:     label,
		Status:    status,
		StartDate: startDate,
		EndDate:   endDate,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

type TodoTaskTitle struct {
	ID     int    `validate:"required"`
	Title  string `validate:"required"`
	Label  null.String
	Status string `validate:"required"`
}

func NewTodoTaskTitle(
	id int,
	title string,
	label null.String,
	status string,
) *TodoTaskTitle {
	return &TodoTaskTitle{
		ID:     id,
		Title:  title,
		Label:  label,
		Status: status,
	}
}
