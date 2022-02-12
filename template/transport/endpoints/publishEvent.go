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

// Endpoints exposes all endpoints.
type Endpoints struct {
	HellowEndpoint pkg.Endpoint
	CreateEndpoint pkg.Endpoint
}

// MakeEndpoints takes service and returns Endpoints
func MakeEndpoints(tracer trace.Tracer, u user.Service) Endpoints {
	return Endpoints{
		HellowEndpoint: helloEndpointHandler(tracer),
		CreateEndpoint: createEndpointHandler(tracer, u),
	}
}

// createNotifHandler to recv email from http as json send the pubAck
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
