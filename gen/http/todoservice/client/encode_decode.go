// Code generated by goa v3.12.1, DO NOT EDIT.
//
// todoservice HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/TsubasaKanemitsu/golang-todo-app/design

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	todoservice "github.com/TsubasaKanemitsu/golang-todo-app/gen/todoservice"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// BuildAddTodoTaskRequest instantiates a HTTP request object with method and
// path set to call the "todoservice" service "addTodoTask" endpoint
func (c *Client) BuildAddTodoTaskRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: AddTodoTaskTodoservicePath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("todoservice", "addTodoTask", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeAddTodoTaskRequest returns an encoder for requests sent to the
// todoservice addTodoTask server.
func EncodeAddTodoTaskRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*todoservice.AddTodoTaskPayload)
		if !ok {
			return goahttp.ErrInvalidType("todoservice", "addTodoTask", "*todoservice.AddTodoTaskPayload", v)
		}
		body := NewAddTodoTaskRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("todoservice", "addTodoTask", err)
		}
		return nil
	}
}

// DecodeAddTodoTaskResponse returns a decoder for responses returned by the
// todoservice addTodoTask endpoint. restoreBody controls whether the response
// body should be restored after having been read.
func DecodeAddTodoTaskResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body bool
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("todoservice", "addTodoTask", err)
			}
			return body, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("todoservice", "addTodoTask", resp.StatusCode, string(body))
		}
	}
}

// BuildGetTodoTaskRequest instantiates a HTTP request object with method and
// path set to call the "todoservice" service "GetTodoTask" endpoint
func (c *Client) BuildGetTodoTaskRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id int
	)
	{
		p, ok := v.(*todoservice.GetTodoTaskPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("todoservice", "GetTodoTask", "*todoservice.GetTodoTaskPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetTodoTaskTodoservicePath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("todoservice", "GetTodoTask", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeGetTodoTaskResponse returns a decoder for responses returned by the
// todoservice GetTodoTask endpoint. restoreBody controls whether the response
// body should be restored after having been read.
func DecodeGetTodoTaskResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetTodoTaskResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("todoservice", "GetTodoTask", err)
			}
			err = ValidateGetTodoTaskResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("todoservice", "GetTodoTask", err)
			}
			res := NewGetTodoTaskTodoTaskInfoOK(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("todoservice", "GetTodoTask", resp.StatusCode, string(body))
		}
	}
}

// BuildGetTodoTaskListRequest instantiates a HTTP request object with method
// and path set to call the "todoservice" service "GetTodoTaskList" endpoint
func (c *Client) BuildGetTodoTaskListRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetTodoTaskListTodoservicePath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("todoservice", "GetTodoTaskList", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeGetTodoTaskListResponse returns a decoder for responses returned by
// the todoservice GetTodoTaskList endpoint. restoreBody controls whether the
// response body should be restored after having been read.
func DecodeGetTodoTaskListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetTodoTaskListResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("todoservice", "GetTodoTaskList", err)
			}
			for _, e := range body {
				if e != nil {
					if err2 := ValidateTodoTaskTitleListResponse(e); err2 != nil {
						err = goa.MergeErrors(err, err2)
					}
				}
			}
			if err != nil {
				return nil, goahttp.ErrValidationError("todoservice", "GetTodoTaskList", err)
			}
			res := NewGetTodoTaskListTodoTaskTitleListOK(body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("todoservice", "GetTodoTaskList", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateTodoTaskRequest instantiates a HTTP request object with method
// and path set to call the "todoservice" service "UpdateTodoTask" endpoint
func (c *Client) BuildUpdateTodoTaskRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id int
	)
	{
		p, ok := v.(*todoservice.UpdateTodoTaskPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("todoservice", "UpdateTodoTask", "*todoservice.UpdateTodoTaskPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateTodoTaskTodoservicePath(id)}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("todoservice", "UpdateTodoTask", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeUpdateTodoTaskResponse returns a decoder for responses returned by the
// todoservice UpdateTodoTask endpoint. restoreBody controls whether the
// response body should be restored after having been read.
func DecodeUpdateTodoTaskResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body bool
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("todoservice", "UpdateTodoTask", err)
			}
			return body, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("todoservice", "UpdateTodoTask", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteTodoTaskRequest instantiates a HTTP request object with method
// and path set to call the "todoservice" service "DeleteTodoTask" endpoint
func (c *Client) BuildDeleteTodoTaskRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id int
	)
	{
		p, ok := v.(*todoservice.DeleteTodoTaskPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("todoservice", "DeleteTodoTask", "*todoservice.DeleteTodoTaskPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteTodoTaskTodoservicePath(id)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("todoservice", "DeleteTodoTask", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeDeleteTodoTaskResponse returns a decoder for responses returned by the
// todoservice DeleteTodoTask endpoint. restoreBody controls whether the
// response body should be restored after having been read.
func DecodeDeleteTodoTaskResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body bool
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("todoservice", "DeleteTodoTask", err)
			}
			return body, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("todoservice", "DeleteTodoTask", resp.StatusCode, string(body))
		}
	}
}

// unmarshalTodoTaskTitleListResponseToTodoserviceTodoTaskTitleList builds a
// value of type *todoservice.TodoTaskTitleList from a value of type
// *TodoTaskTitleListResponse.
func unmarshalTodoTaskTitleListResponseToTodoserviceTodoTaskTitleList(v *TodoTaskTitleListResponse) *todoservice.TodoTaskTitleList {
	res := &todoservice.TodoTaskTitleList{
		ID:     *v.ID,
		Title:  *v.Title,
		Status: *v.Status,
		Label:  v.Label,
	}

	return res
}