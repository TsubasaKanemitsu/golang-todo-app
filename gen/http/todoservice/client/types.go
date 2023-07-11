// Code generated by goa v3.12.1, DO NOT EDIT.
//
// todoservice HTTP client types
//
// Command:
// $ goa gen github.com/TsubasaKanemitsu/golang-todo-app/design

package client

import (
	todoservice "github.com/TsubasaKanemitsu/golang-todo-app/gen/todoservice"
	goa "goa.design/goa/v3/pkg"
)

// AddTodoTaskRequestBody is the type of the "todoservice" service
// "addTodoTask" endpoint HTTP request body.
type AddTodoTaskRequestBody struct {
	// Todoタスクのタイトル
	Title string `form:"title" json:"title" xml:"title"`
	// Todoタスクの説明
	Contents *string `form:"contents,omitempty" json:"contents,omitempty" xml:"contents,omitempty"`
	// Todoタスクのラベル
	Label *string `form:"label,omitempty" json:"label,omitempty" xml:"label,omitempty"`
	// タスクを割り当てられた人の名前
	Asignee *string `form:"asignee,omitempty" json:"asignee,omitempty" xml:"asignee,omitempty"`
	// Todoタスクの開始予定日
	StartDate string `form:"start_date" json:"start_date" xml:"start_date"`
	// Todoタスクの終了予定日
	EndDate string `form:"end_date" json:"end_date" xml:"end_date"`
}

// GetTodoTaskResponseBody is the type of the "todoservice" service
// "GetTodoTask" endpoint HTTP response body.
type GetTodoTaskResponseBody struct {
	// TodoタスクのユニークID
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Todoタスクのタイトル
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Todoタスクの説明
	Contents *string `form:"contents,omitempty" json:"contents,omitempty" xml:"contents,omitempty"`
	// Todoタスクの進捗状況
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Todoタスクのラベル
	Label *string `form:"label,omitempty" json:"label,omitempty" xml:"label,omitempty"`
	// タスクを割り当てられた人の名前
	Asignee *string `form:"asignee,omitempty" json:"asignee,omitempty" xml:"asignee,omitempty"`
	// Todoタスクの開始予定日
	StartDate *string `form:"start_date,omitempty" json:"start_date,omitempty" xml:"start_date,omitempty"`
	// Todoタスクの終了予定日
	EndDate *string `form:"end_date,omitempty" json:"end_date,omitempty" xml:"end_date,omitempty"`
}

// GetTodoTaskListResponseBody is the type of the "todoservice" service
// "GetTodoTaskList" endpoint HTTP response body.
type GetTodoTaskListResponseBody []*TodoTaskTitleListResponse

// TodoTaskTitleListResponse is used to define fields on response body types.
type TodoTaskTitleListResponse struct {
	// TodoタスクのユニークID
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Todoタスクのタイトル
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Todoタスクの進捗状況
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Todoタスクのラベル
	Label *string `form:"label,omitempty" json:"label,omitempty" xml:"label,omitempty"`
}

// NewAddTodoTaskRequestBody builds the HTTP request body from the payload of
// the "addTodoTask" endpoint of the "todoservice" service.
func NewAddTodoTaskRequestBody(p *todoservice.AddTodoTaskPayload) *AddTodoTaskRequestBody {
	body := &AddTodoTaskRequestBody{
		Title:     p.Title,
		Contents:  p.Contents,
		Label:     p.Label,
		Asignee:   p.Asignee,
		StartDate: p.StartDate,
		EndDate:   p.EndDate,
	}
	return body
}

// NewGetTodoTaskTodoTaskInfoOK builds a "todoservice" service "GetTodoTask"
// endpoint result from a HTTP "OK" response.
func NewGetTodoTaskTodoTaskInfoOK(body *GetTodoTaskResponseBody) *todoservice.TodoTaskInfo {
	v := &todoservice.TodoTaskInfo{
		ID:        *body.ID,
		Title:     *body.Title,
		Contents:  body.Contents,
		Status:    *body.Status,
		Label:     body.Label,
		Asignee:   body.Asignee,
		StartDate: body.StartDate,
		EndDate:   body.EndDate,
	}

	return v
}

// NewGetTodoTaskListTodoTaskTitleListOK builds a "todoservice" service
// "GetTodoTaskList" endpoint result from a HTTP "OK" response.
func NewGetTodoTaskListTodoTaskTitleListOK(body []*TodoTaskTitleListResponse) []*todoservice.TodoTaskTitleList {
	v := make([]*todoservice.TodoTaskTitleList, len(body))
	for i, val := range body {
		v[i] = unmarshalTodoTaskTitleListResponseToTodoserviceTodoTaskTitleList(val)
	}

	return v
}

// ValidateGetTodoTaskResponseBody runs the validations defined on
// GetTodoTaskResponseBody
func ValidateGetTodoTaskResponseBody(body *GetTodoTaskResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "body"))
	}
	if body.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "body"))
	}
	if body.StartDate != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.start_date", *body.StartDate, goa.FormatDate))
	}
	if body.EndDate != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.end_date", *body.EndDate, goa.FormatDate))
	}
	return
}

// ValidateTodoTaskTitleListResponse runs the validations defined on
// todoTaskTitleListResponse
func ValidateTodoTaskTitleListResponse(body *TodoTaskTitleListResponse) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "body"))
	}
	if body.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "body"))
	}
	return
}