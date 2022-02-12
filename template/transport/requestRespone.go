package transport

// GenericResponse represents a common response of APIs.
type GenericResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
