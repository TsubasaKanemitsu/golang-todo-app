// Code generated by goa v3.12.1, DO NOT EDIT.
//
// service HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/TsubasaKanemitsu/golang-todo-app/backend/design

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	service "github.com/TsubasaKanemitsu/golang-todo-app/gen/service"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// BuildAddTodoTaskRequest instantiates a HTTP request object with method and
// path set to call the "service" service "addTodoTask" endpoint
func (c *Client) BuildAddTodoTaskRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: AddTodoTaskServicePath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("service", "addTodoTask", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeAddTodoTaskRequest returns an encoder for requests sent to the service
// addTodoTask server.
func EncodeAddTodoTaskRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*service.AddTodoTaskPayload)
		if !ok {
			return goahttp.ErrInvalidType("service", "addTodoTask", "*service.AddTodoTaskPayload", v)
		}
		body := NewAddTodoTaskRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("service", "addTodoTask", err)
		}
		return nil
	}
}

// DecodeAddTodoTaskResponse returns a decoder for responses returned by the
// service addTodoTask endpoint. restoreBody controls whether the response body
// should be restored after having been read.
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
				return nil, goahttp.ErrDecodingError("service", "addTodoTask", err)
			}
			return body, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("service", "addTodoTask", resp.StatusCode, string(body))
		}
	}
}

// BuildGetTodoTaskRequest instantiates a HTTP request object with method and
// path set to call the "service" service "GetTodoTask" endpoint
func (c *Client) BuildGetTodoTaskRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id int
	)
	{
		p, ok := v.(*service.GetTodoTaskPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("service", "GetTodoTask", "*service.GetTodoTaskPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetTodoTaskServicePath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("service", "GetTodoTask", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeGetTodoTaskResponse returns a decoder for responses returned by the
// service GetTodoTask endpoint. restoreBody controls whether the response body
// should be restored after having been read.
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
				return nil, goahttp.ErrDecodingError("service", "GetTodoTask", err)
			}
			err = ValidateGetTodoTaskResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "GetTodoTask", err)
			}
			res := NewGetTodoTaskTodoTaskInfoOK(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("service", "GetTodoTask", resp.StatusCode, string(body))
		}
	}
}

// BuildGetTodoTaskListRequest instantiates a HTTP request object with method
// and path set to call the "service" service "GetTodoTaskList" endpoint
func (c *Client) BuildGetTodoTaskListRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetTodoTaskListServicePath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("service", "GetTodoTaskList", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeGetTodoTaskListResponse returns a decoder for responses returned by
// the service GetTodoTaskList endpoint. restoreBody controls whether the
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
				return nil, goahttp.ErrDecodingError("service", "GetTodoTaskList", err)
			}
			for _, e := range body {
				if e != nil {
					if err2 := ValidateTodoTaskTitleListResponse(e); err2 != nil {
						err = goa.MergeErrors(err, err2)
					}
				}
			}
			if err != nil {
				return nil, goahttp.ErrValidationError("service", "GetTodoTaskList", err)
			}
			res := NewGetTodoTaskListTodoTaskTitleListOK(body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("service", "GetTodoTaskList", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateTodoTaskRequest instantiates a HTTP request object with method
// and path set to call the "service" service "UpdateTodoTask" endpoint
func (c *Client) BuildUpdateTodoTaskRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id int
	)
	{
		p, ok := v.(*service.UpdateTodoTaskPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("service", "UpdateTodoTask", "*service.UpdateTodoTaskPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateTodoTaskServicePath(id)}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("service", "UpdateTodoTask", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeUpdateTodoTaskResponse returns a decoder for responses returned by the
// service UpdateTodoTask endpoint. restoreBody controls whether the response
// body should be restored after having been read.
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
				return nil, goahttp.ErrDecodingError("service", "UpdateTodoTask", err)
			}
			return body, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("service", "UpdateTodoTask", resp.StatusCode, string(body))
		}
	}
}

// BuildDELETETodoTaskRequest instantiates a HTTP request object with method
// and path set to call the "service" service "DELETETodoTask" endpoint
func (c *Client) BuildDELETETodoTaskRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id int
	)
	{
		p, ok := v.(*service.DELETETodoTaskPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("service", "DELETETodoTask", "*service.DELETETodoTaskPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DELETETodoTaskServicePath(id)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("service", "DELETETodoTask", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeDELETETodoTaskResponse returns a decoder for responses returned by the
// service DELETETodoTask endpoint. restoreBody controls whether the response
// body should be restored after having been read.
func DecodeDELETETodoTaskResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				return nil, goahttp.ErrDecodingError("service", "DELETETodoTask", err)
			}
			return body, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("service", "DELETETodoTask", resp.StatusCode, string(body))
		}
	}
}

// unmarshalTodoTaskTitleListResponseToServiceTodoTaskTitleList builds a value
// of type *service.TodoTaskTitleList from a value of type
// *TodoTaskTitleListResponse.
func unmarshalTodoTaskTitleListResponseToServiceTodoTaskTitleList(v *TodoTaskTitleListResponse) *service.TodoTaskTitleList {
	res := &service.TodoTaskTitleList{
		ID:    *v.ID,
		Title: *v.Title,
	}

	return res
}