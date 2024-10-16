package errors

import (
	"io"
	"net/http"
)

type HTTPError struct {
	StatusCode int
	Name       string
	Message    string
	Original   interface{}
}

func (err HTTPError) Error() string {
	return err.Message
}

func NewHTTPErrorFromReponse(res *http.Response) HTTPError {
	body := res.Body

	httpErr := HTTPError{}
	switch res.StatusCode {
	case http.StatusNotFound:
		httpErr.StatusCode = res.StatusCode
		httpErr.Name = "not_found_error"
		httpErr.Message = httpErr.Name
	case http.StatusUnauthorized:
		httpErr.StatusCode = res.StatusCode
		httpErr.Name = "unauthorized_error"
		httpErr.Message = httpErr.Name
	case http.StatusForbidden:
		httpErr.StatusCode = res.StatusCode
		httpErr.Name = "forbidden_error"
		httpErr.Message = httpErr.Name
	case http.StatusLocked:
		httpErr.StatusCode = res.StatusCode
		httpErr.Name = "locked_error"
		httpErr.Message = httpErr.Name
	case http.StatusGone:
		httpErr.StatusCode = res.StatusCode
		httpErr.Name = "gone_error"
		httpErr.Message = httpErr.Name
	case http.StatusConflict:
		httpErr.StatusCode = res.StatusCode
		httpErr.Name = "conflict_error"
		httpErr.Message = httpErr.Name
	case http.StatusBadGateway:
		httpErr.StatusCode = res.StatusCode
		httpErr.Name = "bad_gateway_error"
		httpErr.Message = httpErr.Name
	case http.StatusBadRequest:
		httpErr.StatusCode = res.StatusCode
		httpErr.Name = "bad_request_error"
		httpErr.Message = httpErr.Name
	case http.StatusGatewayTimeout:
		httpErr.StatusCode = res.StatusCode
		httpErr.Name = "gateway_timeout_error"
		httpErr.Message = httpErr.Name
	case http.StatusTooManyRequests:
		httpErr.StatusCode = res.StatusCode
		httpErr.Name = "too_may_requests_error"
		httpErr.Message = httpErr.Name
	case http.StatusServiceUnavailable:
		httpErr.StatusCode = res.StatusCode
		httpErr.Name = "service_unavailable_error"
		httpErr.Message = httpErr.Name
	default:
		httpErr.StatusCode = http.StatusInternalServerError
		httpErr.Name = "internal_server_error"
		httpErr.Message = httpErr.Name
	}

	b, err := io.ReadAll(body)

	if err != nil {
		return httpErr
	}

	original := string(b)
	httpErr.Original = original

	return httpErr
}
