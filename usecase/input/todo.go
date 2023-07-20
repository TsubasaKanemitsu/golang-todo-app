package input

import (
	"fmt"

	"github.com/TsubasaKanemitsu/golang-todo-app/domain/model"
	"github.com/TsubasaKanemitsu/golang-todo-app/gen/todoservice"
	"github.com/volatiletech/null/v8"
)

func ToMTodoTaskInfo(p *todoservice.AddTodoTaskPayload) (*model.TodoTaskInfo, error) {
	contents := null.StringFromPtr(p.Contents)
	assignee := null.StringFromPtr(p.Asignee)
	label := null.StringFromPtr(p.Label)
	// status := null.StringFromPtr(p.Status)

	m := &model.TodoTaskInfo{
		Title:    p.Title,
		Contents: contents,
		Assignee: assignee,
		Label:    label,
		Status:   p.Status,
	}
	// TODO:
	// Add startDate and endDate
	fmt.Println("m: ", m)
	fmt.Println(p.StartDate)
	return m, nil
}
