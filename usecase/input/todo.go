package input

import (
	"fmt"
	"time"

	"github.com/TsubasaKanemitsu/golang-todo-app/domain/model"
	"github.com/TsubasaKanemitsu/golang-todo-app/gen/todoservice"
	"github.com/volatiletech/null/v8"
)

func ToMTodoTaskInfoForAdd(p *todoservice.AddTodoTaskPayload) (*model.TodoTaskInfo, error) {
	contents := null.StringFromPtr(p.Contents)
	assignee := null.StringFromPtr(p.Asignee)
	label := null.StringFromPtr(p.Label)

	m := &model.TodoTaskInfo{
		Title:    p.Title,
		Contents: contents,
		Assignee: assignee,
		Label:    label,
		Status:   p.Status,
	}

	// Add StartDate and EndDate
	startDate, err := time.Parse("2006-01-02", p.StartDate)
	if err != nil {
		return nil, fmt.Errorf("error: time.Parse(), startDate = %v", p.StartDate)
	}
	endDate, err := time.Parse("2006-01-02", p.EndDate)
	if err != nil {
		return nil, fmt.Errorf("error: time.Parse(), endDate = %v", p.StartDate)
	}
	m.StartDate = startDate
	m.EndDate = endDate

	return m, nil
}

func ToMTodoTaskInfoForUpdate(p *todoservice.UpdateTodoTaskPayload) (*model.TodoTaskInfo, error) {
	contents := null.StringFromPtr(p.Contents)
	assignee := null.StringFromPtr(p.Asignee)
	label := null.StringFromPtr(p.Label)

	m := &model.TodoTaskInfo{
		ID:       p.ID,
		Title:    p.Title,
		Contents: contents,
		Assignee: assignee,
		Label:    label,
		Status:   p.Status,
	}

	// Add StartDate and EndDate
	startDate, err := time.Parse("2006-01-02", p.StartDate)
	if err != nil {
		return nil, fmt.Errorf("error: time.Parse(), startDate = %v", p.StartDate)
	}
	endDate, err := time.Parse("2006-01-02", p.EndDate)
	if err != nil {
		return nil, fmt.Errorf("error: time.Parse(), endDate = %v", p.StartDate)
	}
	m.StartDate = startDate
	m.EndDate = endDate

	return m, nil
}
