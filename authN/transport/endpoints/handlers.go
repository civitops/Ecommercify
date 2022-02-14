package endpoints

import (
	"context"
	"encoding/json"
	"github.com/civitops/Ecommercify/authN/pkg"
	"github.com/civitops/Ecommercify/authN/transport"
	"go.opentelemetry.io/otel/trace"
	"io"
	"io/ioutil"
)

// Endpoints exposes all endpoints.
type Endpoints struct {
	HellowEndpoint pkg.Endpoint
}

// MakeEndpoints takes service and returns Endpoints
func MakeEndpoints(tracer trace.Tracer) Endpoints {
	return Endpoints{
		HellowEndpoint: helloEndpointHandler(tracer),
	}
}

// createNotifHandler to recv email from http as json send the pubAck
func helloEndpointHandler(tracer trace.Tracer) pkg.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_, span := tracer.Start(ctx, "Hello-endpoint-handler")
		defer span.End()

		var response transport.GenericResponse
		var body map[string]interface{}

		data, err := ioutil.ReadAll(request.(io.Reader))
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(data, &body)
		if err != nil {
			return response, err
		}

		response.Message = "ok"
		response.Data = body
		return response, nil
	}
}
