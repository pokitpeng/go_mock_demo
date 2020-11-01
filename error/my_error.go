package error

import (
	"errors"
	"fmt"
)

// MyError define my error
type MyError struct {
	err  error // err
	code int   // err code
}

// New create new error
func New(text string) error {
	return errors.New(text)
}

// NewWithCode create new error
func NewWithCode(code int) error {
	return MyError{err: errors.New(Trans(code)), code: code}
}

// Wrap create new error
func Wrap(err error, code int) error {
	return MyError{
		err:  err,
		code: code,
	}
}

func (e MyError) Error() string {
	return e.err.Error()
}

func (e MyError) String() string {
	return fmt.Sprintf("code:%d,err:%s", e.code, e.err.Error())
}

// Code get error code
func (e MyError) Code() int {
	return e.code
}
