package entity

import (
	"github.com/palantir/stacktrace"
)

// Code - Max value is 4294967295
type Code = stacktrace.ErrorCode

// ErrCode extracts the error code from an error.
var ErrCode = stacktrace.GetCode

// New is a drop-in replacement for fmt.Errorf that includes line number information.
var New = stacktrace.NewError

// NewWithCode is similar to New but also attaches an error code.
var NewWithCode = stacktrace.NewErrorWithCode

// RootCause unwraps the original error that caused the current one.
var RootCause = stacktrace.RootCause

// Wrap an error to include line number information.
var Wrap = stacktrace.Propagate

// WrapWithCode is similar to Wrap but also attaches an error code.
var WrapWithCode = stacktrace.PropagateWithCode

// Wrapf is similar to Wrap but the msg and vals arguments work like the ones for fmt.Errorf.
var Wrapf = stacktrace.Propagate

type Message struct {
	StatusCode    int    `json:"status_code"`
	EN            string `json:"en"`
	ID            string `json:"id"`
	HasAnnotation bool
	CustomMessage bool
}

// ErrorMessage - Mapping Error Code as Human Message
type ErrorMessage map[Code]Message
