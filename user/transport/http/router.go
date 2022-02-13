package http

import (
	"net/http"

	"github.com/civitops/Ecommercify/user/pkg"
	"github.com/civitops/Ecommercify/user/transport/endpoints"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type httpSvc struct {
	t   trace.Tracer
	log *zap.SugaredLogger
}

// NewHTTPService takes all the endpoints and returns handler.
func NewHTTPService(endpoints endpoints.Endpoints, t trace.Tracer, l *zap.SugaredLogger) http.Handler {
	h := httpSvc{
		t:   t,
		log: l,
	}

	r := echo.New()

	r.Use(middleware.Recover())
	r.Use(middleware.Logger())

	user := r.Group("/user-svc/v1")
	{
		user.POST("/hello", h.endpointRequestEncoder(endpoints.HellowEndpoint))
		user.POST("/create", h.endpointRequestEncoder(endpoints.CreateEndpoint))
		user.PATCH("/update", h.endpointRequestEncoder(endpoints.UpdateEndpoint))
	}

	return r
}

// endpointRequestEncoder encodes request and does error handling
// and send response.
func (h *httpSvc) endpointRequestEncoder(endpoint pkg.Endpoint) echo.HandlerFunc {
	return func(c echo.Context) error {
		var statusCode int

		ctx, span := h.t.Start(c.Request().Context(), "endpoint-Req-Encoder")
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
