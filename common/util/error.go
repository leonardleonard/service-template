package util

import (
	"fmt"
)

type Error interface {
	error

	ErrCode() int
	ErrMessage() string
}

type CommonError struct {
	Code       int
	HttpStatus int
	Message    string

	source string
	err    error
}

func (e CommonError) WithError(err error) CommonError {
	e.err = err
	return e
}

func (e CommonError) WithSource(source string, params ...interface{}) CommonError {
	if len(params) == 0 {
		e.source = source
	} else {
		e.source = fmt.Sprintf(source, params...)
	}

	return e
}

/********** 实现Error接口 **********/

func (e CommonError) ErrCode() int {
	return e.Code
}

func (e CommonError) ErrMessage() string {
	return e.Message
}

func (e CommonError) Error() string {
	if e.err != nil {
		return e.err.Error()
	}

	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}
