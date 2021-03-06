package errors

import (
	"net/http"

	"github.com/go-chi/render"
)

// ErrorResponse represents custom error response with statusText and error description
type ErrorResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

// Render realization method for render.renderer
func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// ErrorInternalServer returns error response with status=500 and given error
func ErrorInternalServer(err error) *ErrorResponse {
	return &ErrorResponse{
		Err:            err,
		HTTPStatusCode: 500, //nolint
		StatusText:     "Internal Server Error",
		ErrorText:      err.Error(),
	}
}

// ErrorInvalidRequest return error response with status = 400 and given error
func ErrorInvalidRequest(err error) *ErrorResponse {
	return &ErrorResponse{
		Err:            err,
		HTTPStatusCode: 400, //nolint
		StatusText:     "Invalid request",
		ErrorText:      err.Error(),
	}
}

// ErrorNotFound return 404 with given error text
func ErrorNotFound(err error) *ErrorResponse {
	return &ErrorResponse{
		HTTPStatusCode: 404, //nolint
		StatusText:     "Resource not found",
		ErrorText:      err.Error(),
	}
}

// ErrorForbidden return 403 with given error text
func ErrorForbidden(err error) *ErrorResponse {
	return &ErrorResponse{
		HTTPStatusCode: 403, //nolint
		StatusText:     "Forbidden",
		ErrorText:      err.Error(),
	}
}

// ErrorUnauthorized return 401 with given error text
func ErrorUnauthorized(err error) *ErrorResponse {
	return &ErrorResponse{
		HTTPStatusCode: 401, //nolint
		StatusText:     "Unauthorized",
		ErrorText:      err.Error(),
	}
}

// ErrorNotAcceptable return 406 with given error text
func ErrorNotAcceptable(err error) *ErrorResponse {
	return &ErrorResponse{
		HTTPStatusCode: 406, //nolint
		StatusText:     "Error",
		ErrorText:      err.Error(),
	}
}

// ErrNotFound is default router page not found
var ErrNotFound = &ErrorResponse{HTTPStatusCode: 404, StatusText: "Page not found."} //nolint

// ErrMethodNotAllowed is default 405 router method not allowed
var ErrMethodNotAllowed = &ErrorResponse{HTTPStatusCode: 405, StatusText: "Method not allowed."} //nolint
