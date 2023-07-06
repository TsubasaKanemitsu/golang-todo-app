// Code generated by goa v3.12.1, DO NOT EDIT.
//
// service HTTP server types
//
// Command:
// $ goa gen github.com/TsubasaKanemitsu/golang-todo-app/backend/design

package server

import (
	service "github.com/TsubasaKanemitsu/golang-todo-app/gen/service"
	goa "goa.design/goa/v3/pkg"
)

// AddTodoTaskRequestBody is the type of the "service" service "addTodoTask"
// endpoint HTTP request body.
type AddTodoTaskRequestBody struct {
	// Todoタスクのタイトル
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Todoタスクの説明
	Contents *string `form:"contents,omitempty" json:"contents,omitempty" xml:"contents,omitempty"`
	// Todoタスクのラベル
	Label *string `form:"label,omitempty" json:"label,omitempty" xml:"label,omitempty"`
	// タスクを割り当てられた人の名前
	Asignee *string `form:"asignee,omitempty" json:"asignee,omitempty" xml:"asignee,omitempty"`
	// Todoタスクの開始予定日
	StartDate *string `form:"start_date,omitempty" json:"start_date,omitempty" xml:"start_date,omitempty"`
	// Todoタスクの終了予定日
	EndDate *string `form:"end_date,omitempty" json:"end_date,omitempty" xml:"end_date,omitempty"`
}

// GetTodoTaskResponseBody is the type of the "service" service "GetTodoTask"
// endpoint HTTP response body.
type GetTodoTaskResponseBody struct {
	// TodoタスクのユニークID
	ID int `form:"id" json:"id" xml:"id"`
	// Todoタスクのタイトル
	Title string `form:"title" json:"title" xml:"title"`
	// Todoタスクの説明
	Contents *string `form:"contents,omitempty" json:"contents,omitempty" xml:"contents,omitempty"`
	// Todoタスクの進捗状況
	Status string `form:"status" json:"status" xml:"status"`
	// Todoタスクのラベル
	Label *string `form:"label,omitempty" json:"label,omitempty" xml:"label,omitempty"`
	// タスクを割り当てられた人の名前
	Asignee *string `form:"asignee,omitempty" json:"asignee,omitempty" xml:"asignee,omitempty"`
	// Todoタスクの開始予定日
	StartDate *string `form:"start_date,omitempty" json:"start_date,omitempty" xml:"start_date,omitempty"`
	// Todoタスクの終了予定日
	EndDate *string `form:"end_date,omitempty" json:"end_date,omitempty" xml:"end_date,omitempty"`
}

// GetTodoTaskListResponseBody is the type of the "service" service
// "GetTodoTaskList" endpoint HTTP response body.
type GetTodoTaskListResponseBody []*TodoTaskTitleListResponse

// TodoTaskTitleListResponse is used to define fields on response body types.
type TodoTaskTitleListResponse struct {
	// TodoタスクのユニークID
	ID int `form:"id" json:"id" xml:"id"`
	// Todoタスクのタイトル
	Title string `form:"title" json:"title" xml:"title"`
	// Todoタスクの進捗状況
	Status string `form:"status" json:"status" xml:"status"`
	// Todoタスクのラベル
	Label *string `form:"label,omitempty" json:"label,omitempty" xml:"label,omitempty"`
}

// NewGetTodoTaskResponseBody builds the HTTP response body from the result of
// the "GetTodoTask" endpoint of the "service" service.
func NewGetTodoTaskResponseBody(res *service.TodoTaskInfo) *GetTodoTaskResponseBody {
	body := &GetTodoTaskResponseBody{
		ID:        res.ID,
		Title:     res.Title,
		Contents:  res.Contents,
		Status:    res.Status,
		Label:     res.Label,
		Asignee:   res.Asignee,
		StartDate: res.StartDate,
		EndDate:   res.EndDate,
	}
	return body
}

// NewGetTodoTaskListResponseBody builds the HTTP response body from the result
// of the "GetTodoTaskList" endpoint of the "service" service.
func NewGetTodoTaskListResponseBody(res []*service.TodoTaskTitleList) GetTodoTaskListResponseBody {
	body := make([]*TodoTaskTitleListResponse, len(res))
	for i, val := range res {
		body[i] = marshalServiceTodoTaskTitleListToTodoTaskTitleListResponse(val)
	}
	return body
}

// NewAddTodoTaskPayload builds a service service addTodoTask endpoint payload.
func NewAddTodoTaskPayload(body *AddTodoTaskRequestBody) *service.AddTodoTaskPayload {
	v := &service.AddTodoTaskPayload{
		Title:     *body.Title,
		Contents:  body.Contents,
		Label:     body.Label,
		Asignee:   body.Asignee,
		StartDate: *body.StartDate,
		EndDate:   *body.EndDate,
	}

	return v
}

// NewGetTodoTaskPayload builds a service service GetTodoTask endpoint payload.
func NewGetTodoTaskPayload(id int) *service.GetTodoTaskPayload {
	v := &service.GetTodoTaskPayload{}
	v.ID = id

	return v
}

// NewUpdateTodoTaskPayload builds a service service UpdateTodoTask endpoint
// payload.
func NewUpdateTodoTaskPayload(id int) *service.UpdateTodoTaskPayload {
	v := &service.UpdateTodoTaskPayload{}
	v.ID = id

	return v
}

// NewDELETETodoTaskPayload builds a service service DELETETodoTask endpoint
// payload.
func NewDELETETodoTaskPayload(id int) *service.DELETETodoTaskPayload {
	v := &service.DELETETodoTaskPayload{}
	v.ID = id

	return v
}

// ValidateAddTodoTaskRequestBody runs the validations defined on
// AddTodoTaskRequestBody
func ValidateAddTodoTaskRequestBody(body *AddTodoTaskRequestBody) (err error) {
	if body.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "body"))
	}
	if body.StartDate == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("start_date", "body"))
	}
	if body.EndDate == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("end_date", "body"))
	}
	if body.StartDate != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.start_date", *body.StartDate, goa.FormatDate))
	}
	if body.EndDate != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.end_date", *body.EndDate, goa.FormatDate))
	}
	return
}
