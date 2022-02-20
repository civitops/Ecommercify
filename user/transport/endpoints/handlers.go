package endpoints

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/civitops/Ecommercify/user/implementation/user"
	"github.com/civitops/Ecommercify/user/pkg"
	"github.com/civitops/Ecommercify/user/transport"
	"go.opentelemetry.io/otel/trace"
)

type Deleterequest struct {
	ID uint `json:"id"`
}
type GetRequest struct {
	Select string                 `json:"sel"`
	Where  map[string]interface{} `json:"where"`
}

// Endpoints exposes all endpoints.
type Endpoints struct {
	HellowEndpoint pkg.Endpoint
	CreateEndpoint pkg.Endpoint
	UpdateEndpoint pkg.Endpoint
	DeleteEndpoint pkg.Endpoint
	GetEndpoint    pkg.Endpoint
}

// MakeEndpoints takes service and returns Endpoints
func MakeEndpoints(tracer trace.Tracer, u user.Service) Endpoints {
	return Endpoints{
		HellowEndpoint: helloEndpointHandler(tracer),
		CreateEndpoint: createEndpointHandler(tracer, u),
		UpdateEndpoint: updateEndpointHandler(tracer, u),
		DeleteEndpoint: deleteEndpointHandler(tracer, u),
		GetEndpoint:    getEndpointHandler(tracer, u),
	}
}

// helloEndpointHandler for testing http service
func helloEndpointHandler(tracer trace.Tracer) pkg.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_, span := tracer.Start(ctx, "Hello-endpoint-handler")
		defer span.End()

		var response transport.GenericResponse
		var body user.Entity

		data, err := ioutil.ReadAll(request.(io.Reader))
		if err != nil {
			return nil, pkg.UserErr{
				Code: http.StatusBadRequest,
				Err:  err,
			}
		}

		err = json.Unmarshal(data, &body)
		if err != nil {
			return nil, pkg.UserErr{
				Code: http.StatusBadRequest,
				Err:  err,
			}
		}

		response.Message = "ok"
		response.Data = body
		return response, nil
	}
}

// createUserHandler to recv email from http as json send the pubAck
func createEndpointHandler(tracer trace.Tracer, u user.Service) pkg.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		ctxSpan, span := tracer.Start(ctx, "create-endpoint-handler")
		defer span.End()

		var response transport.GenericResponse
		var body user.Entity

		data, err := ioutil.ReadAll(request.(io.Reader))
		if err != nil {
			return nil, pkg.UserErr{
				Code: http.StatusBadRequest,
				Err:  err,
			}
		}

		err = json.Unmarshal(data, &body)
		if err != nil {
			return nil, pkg.UserErr{
				Code: http.StatusBadRequest,
				Err:  err,
			}
		}

		insertedID, err := u.Create(ctxSpan, body)
		if err != nil {
			return nil, pkg.UserErr{
				Code: http.StatusInternalServerError,
				Err:  err,
			}
		}

		response.Message = "ok"
		response.Data = insertedID
		return response, nil
	}
}

func updateEndpointHandler(tracer trace.Tracer, u user.Service) pkg.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		ctxSpan, span := tracer.Start(ctx, "update-endpoint-handler")
		defer span.End()

		var response transport.GenericResponse
		var body user.Entity

		data, err := ioutil.ReadAll(request.(io.Reader))
		if err != nil {
			return nil, pkg.UserErr{
				Code: http.StatusBadRequest,
				Err:  err,
			}
		}

		err = json.Unmarshal(data, &body)
		if err != nil {
			return nil, pkg.UserErr{
				Code: http.StatusBadRequest,
				Err:  err,
			}
		}

		err = u.Update(ctxSpan, body)
		if err != nil {
			return nil, pkg.UserErr{
				Code: http.StatusInternalServerError,
				Err:  err,
			}
		}

		response.Message = "ok"
		return response, nil
	}
}

func deleteEndpointHandler(tracer trace.Tracer, u user.Service) pkg.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		ctxSpan, span := tracer.Start(ctx, "delete-endpoint-handler")
		defer span.End()

		var response transport.GenericResponse
		var body Deleterequest

		data, err := ioutil.ReadAll(request.(io.Reader))
		if err != nil {
			return nil, pkg.UserErr{
				Code: http.StatusBadRequest,
				Err:  err,
			}
		}

		err = json.Unmarshal(data, &body)
		if err != nil {
			return nil, pkg.UserErr{
				Code: http.StatusBadRequest,
				Err:  err,
			}
		}
		err = u.Delete(ctxSpan, body.ID)
		if err != nil {
			return nil, pkg.UserErr{
				Code: http.StatusInternalServerError,
				Err:  err,
			}
		}
		response.Message = "ok"
		response.Data = body
		return response, nil
	}
}

func getEndpointHandler(tracer trace.Tracer, u user.Service) pkg.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		ctxSpan, span := tracer.Start(ctx, "get-endpoint-handler")
		defer span.End()

		var response transport.GenericResponse
		var body GetRequest

		data, err := ioutil.ReadAll(request.(io.Reader))
		if err != nil {
			return nil, pkg.UserErr{
				Code: http.StatusBadRequest,
				Err:  err,
			}
		}

		err = json.Unmarshal(data, &body)
		if err != nil {
			return nil, pkg.UserErr{
				Code: http.StatusBadRequest,
				Err:  err,
			}
		}

		userRes, err := u.Get(ctxSpan, body.Select, body.Where)
		if err != nil {
			return nil, pkg.UserErr{
				Code: http.StatusInternalServerError,
				Err:  err,
			}
		}
		response.Message = "ok"
		response.Data = userRes
		return response, nil
	}
}
