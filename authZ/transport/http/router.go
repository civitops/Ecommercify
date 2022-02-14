package http

import (
	"net/http"

	"github.com/civitops/Ecommercify/auth/pkg"
	"github.com/civitops/Ecommercify/auth/transport/endpoints"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// NewHTTPService takes all the endpoints and returns handler.
func NewHTTPService(endpoints endpoints.Endpoints, t trace.Tracer) http.Handler {

	r := echo.New()

	r.Use(middleware.Recover())
	r.Use(middleware.Logger())

	notif := r.Group("/notif-svc/v1")
	{
		notif.POST("/create", endpointRequestEncoder(endpoints.HellowEndpoint, t))
	}

	return r
}

// endpointRequestEncoder encodes request and does error handling
// and send response.
func endpointRequestEncoder(endpoint pkg.Endpoint, t trace.Tracer) echo.HandlerFunc {
	return func(c echo.Context) error {
		var statusCode int

		ctx, span := t.Start(c.Request().Context(), "endpoint-Req-Encoder")
		defer span.End()

		// process the request with its handler
		response, err := endpoint(ctx, c.Request().Body)
		if err != nil {
			// if statusCode is not send then return InternalServerErr
			switch e := err.(type) {
			case pkg.Error:
				statusCode = e.Status()

			default:
				statusCode = http.StatusInternalServerError
			}

			c.JSON(statusCode, map[string]interface{}{
				"error":   true,
				"message": err.Error(),
			})

			return err
		}

		// if err did not occur then return Ok status
		span.SetStatus(codes.Ok, "request proccessed suceessfully")
		c.JSON(http.StatusOK, response)
		return nil
	}
}
