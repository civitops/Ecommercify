package endpoints

import (
	"context"
	"notif/pkg"

	"go.opentelemetry.io/otel/trace"
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
		ctx, span := tracer.Start(ctx, "Hello-endpoint-handler")
		defer span.End()

		return nil, nil
	}
}
