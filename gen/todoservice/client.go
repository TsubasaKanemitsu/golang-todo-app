// Code generated by goa v3.12.1, DO NOT EDIT.
//
// todoservice client
//
// Command:
// $ goa gen github.com/TsubasaKanemitsu/golang-todo-app/design

package todoservice

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "todoservice" service client.
type Client struct {
	AddTodoTaskEndpoint     goa.Endpoint
	GetTodoTaskEndpoint     goa.Endpoint
	GetTodoTaskListEndpoint goa.Endpoint
	UpdateTodoTaskEndpoint  goa.Endpoint
	DeleteTodoTaskEndpoint  goa.Endpoint
}

// NewClient initializes a "todoservice" service client given the endpoints.
func NewClient(addTodoTask, getTodoTask, getTodoTaskList, updateTodoTask, deleteTodoTask goa.Endpoint) *Client {
	return &Client{
		AddTodoTaskEndpoint:     addTodoTask,
		GetTodoTaskEndpoint:     getTodoTask,
		GetTodoTaskListEndpoint: getTodoTaskList,
		UpdateTodoTaskEndpoint:  updateTodoTask,
		DeleteTodoTaskEndpoint:  deleteTodoTask,
	}
}

// AddTodoTask calls the "addTodoTask" endpoint of the "todoservice" service.
func (c *Client) AddTodoTask(ctx context.Context, p *AddTodoTaskPayload) (res bool, err error) {
	var ires any
	ires, err = c.AddTodoTaskEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(bool), nil
}

// GetTodoTask calls the "GetTodoTask" endpoint of the "todoservice" service.
func (c *Client) GetTodoTask(ctx context.Context, p *GetTodoTaskPayload) (res *TodoTaskInfo, err error) {
	var ires any
	ires, err = c.GetTodoTaskEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*TodoTaskInfo), nil
}

// GetTodoTaskList calls the "GetTodoTaskList" endpoint of the "todoservice"
// service.
func (c *Client) GetTodoTaskList(ctx context.Context) (res []*TodoTaskTitle, err error) {
	var ires any
	ires, err = c.GetTodoTaskListEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.([]*TodoTaskTitle), nil
}

// UpdateTodoTask calls the "UpdateTodoTask" endpoint of the "todoservice"
// service.
func (c *Client) UpdateTodoTask(ctx context.Context, p *UpdateTodoTaskPayload) (res bool, err error) {
	var ires any
	ires, err = c.UpdateTodoTaskEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(bool), nil
}

// DeleteTodoTask calls the "DeleteTodoTask" endpoint of the "todoservice"
// service.
func (c *Client) DeleteTodoTask(ctx context.Context, p *DeleteTodoTaskPayload) (res bool, err error) {
	var ires any
	ires, err = c.DeleteTodoTaskEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(bool), nil
}
