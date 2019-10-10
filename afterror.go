package afterror

import (
	"fmt"
	"net/http"
)

type Error struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

func New(err error) *Error {
	if err == nil {
		return nil
	}

	return &Error{Err: err, Message: err.Error()}
}

func (aftErr *Error) print() string {
	if aftErr.Type != "" && aftErr.Message != "" {
		return fmt.Sprintf("%s: %s", aftErr.Type, aftErr.Message)
	} else if aftErr.Message != "" {
		return fmt.Sprintf("%s", aftErr.Message)
	}

	return "Unknown Error"
}

func (aftErr *Error) Error() string {
	return aftErr.print()
}

func (aftErr *Error) Wrap(err error) *Error {
	aftErr.Err = err

	return aftErr
}

func (aftErr *Error) IsTypeEqual(err error) bool {
	if e, ok := err.(*Error); ok {
		return aftErr.Type == e.Type
	}

	return false
}

func BadRequest(errType, errMessage string) *Error {
	err := &Error{
		Code:    http.StatusBadRequest,
		Type:    errType,
		Message: errMessage,
	}

	return err
}

func InternalServer(errType, errMessage string) *Error {
	err := &Error{
		Code:    http.StatusInternalServerError,
		Type:    errType,
		Message: errMessage,
	}

	return err
}

func NotFound(errType, errMessage string) *Error {
	err := &Error{
		Code:    http.StatusNotFound,
		Type:    errType,
		Message: errMessage,
	}

	return err
}

func Unauthorized(errType, errMessage string) *Error {
	err := &Error{
		Code:    http.StatusUnauthorized,
		Type:    errType,
		Message: errMessage,
	}

	return err
}

func Forbidden(errType, errMessage string) *Error {
	err := &Error{
		Code:    http.StatusForbidden,
		Type:    errType,
		Message: errMessage,
	}

	return err
}
