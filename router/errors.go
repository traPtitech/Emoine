package router

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"runtime"
)

type option func(*ErrorResponse)

type ErrorResponse struct {
	ErrorBody `json:"errors"`
}
type ErrorBody struct {
	Message       string `json:"message,omitempty"`
	Specification string `json:"specification,omitempty"`
	errorRuntime  RuntimeCallerStruct
}
type RuntimeCallerStruct struct {
	ProgramCounter uintptr
	SourceFile     string
	Line           int
	ok             bool
}

func message(msg string) option {
	return func(er *ErrorResponse) {
		er.Message = msg
	}
}

func specification(spec string) option {
	return func(er *ErrorResponse) {
		er.Specification = spec
	}
}

func newRuntimeCallerStruct(pc uintptr, file string, line int, ok bool) RuntimeCallerStruct {
	return RuntimeCallerStruct{
		ProgramCounter: pc,
		SourceFile:     file,
		Line:           line,
		ok:             ok,
	}
}

func newHTTPErrorResponse(err error, code int, options ...option) *echo.HTTPError {
	he := &echo.HTTPError{
		Code: code,
	}
	er := new(ErrorResponse)

	for _, o := range options {
		o(er)
	}
	if er.Message == "" {
		er.Message = http.StatusText(code)
	}
	if er.errorRuntime == (RuntimeCallerStruct{}) {
		er.errorRuntime = newRuntimeCallerStruct(runtime.Caller(2))
	}
	he.Message = er
	err = fmt.Errorf("%s: %s", er.errorRuntime, err)
	he.SetInternal(err)
	return he
}

func internalServerError(err error, responses ...option) *echo.HTTPError {
	code := http.StatusInternalServerError
	return newHTTPErrorResponse(err, code, responses...)
}

func unauthorized(err error, responses ...option) *echo.HTTPError {
	return newHTTPErrorResponse(err, http.StatusUnauthorized, responses...)
}
func forbidden(err error, responses ...option) *echo.HTTPError {
	return newHTTPErrorResponse(err, http.StatusForbidden, responses...)
}
