package errors

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type HTTPError struct {
	StatusCode int         `json:"statusCode"` // response status code
	Name       string      `json:"name"`       // name related to the given StatusCode
	Message    string      `json:"message"`    // user defined message
	Original   interface{} `json:"original"`   // original body from an http response
}

// NewHTTPError returns a new initialized HTTPError.
func NewHTTPError(code int, message string, err any) *HTTPError {
	httpErr := new(HTTPError)
	switch code {
	case http.StatusNotFound:
		httpErr.StatusCode = code
		httpErr.Name = "not_found_error"
	case http.StatusUnauthorized:
		httpErr.StatusCode = code
		httpErr.Name = "unauthorized_error"
	case http.StatusForbidden:
		httpErr.StatusCode = code
		httpErr.Name = "forbidden_error"
	case http.StatusLocked:
		httpErr.StatusCode = code
		httpErr.Name = "locked_error"
	case http.StatusGone:
		httpErr.StatusCode = code
		httpErr.Name = "gone_error"
	case http.StatusConflict:
		httpErr.StatusCode = code
		httpErr.Name = "conflict_error"
	case http.StatusBadGateway:
		httpErr.StatusCode = code
		httpErr.Name = "bad_gateway_error"
	case http.StatusBadRequest:
		httpErr.StatusCode = code
		httpErr.Name = "bad_request_error"
	case http.StatusGatewayTimeout:
		httpErr.StatusCode = code
		httpErr.Name = "gateway_timeout_error"
	case http.StatusTooManyRequests:
		httpErr.StatusCode = code
		httpErr.Name = "too_may_requests_error"
	case http.StatusServiceUnavailable:
		httpErr.StatusCode = code
		httpErr.Name = "service_unavailable_error"
	default:
		httpErr.StatusCode = http.StatusInternalServerError
		httpErr.Name = "internal_server_error"
	}

	httpErr.Message = message
	httpErr.Original = err

	return httpErr
}

// Marshall converts an HTTPError to json.
func (e HTTPError) Marshall() ([]byte, error) {
	return json.Marshal(e)
}

// WriteResponse writes to the sender an HTTPError in json format.
func (e HTTPError) WriteResponse(w http.ResponseWriter) error {
	b, err := e.Marshall()
	if err != nil {
		return err
	}

	w.Header().Set("content-type", "application/json")
	w.Header().Set("content-length", fmt.Sprintf("%d", len(b)))
	w.WriteHeader(e.StatusCode)
	_, err = w.Write(b)
	return err
}

// Error returns the HTTPError.Message property.
func (err HTTPError) Error() string {
	return err.Message
}

// NewHTTPErrorFromResponse returns a new HTTPError based in a http.Response.
//
//   - It maps the response body to the HTTPError.Original property.
//   - It maps the response status to the HTTPError.Message.
//   - If it fails to read the response body, the HTTPError.Message will reflect it.
func NewHTTPErrorFromResponse(res *http.Response) *HTTPError {
	body := res.Body
	b, err := io.ReadAll(body)
	if err != nil {
		return NewHTTPError(http.StatusInternalServerError, "unable to read response body", err)
	}

	httpErr := NewHTTPError(
		res.StatusCode,
		strings.Join(strings.Split(res.Status, " ")[1:], " "),
		nil,
	)

	original := string(b)
	httpErr.Original = original

	return httpErr
}
