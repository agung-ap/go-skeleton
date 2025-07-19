package errors

import (
	"fmt"
	"go-skeleton/pkg/errors/cache"
	"go-skeleton/pkg/errors/entity"
	"go-skeleton/pkg/errors/general"
	httperr "go-skeleton/pkg/errors/http"
	"go-skeleton/pkg/errors/sql"
	"net/http"
	"strings"

	"github.com/palantir/stacktrace"
)

var svcError map[ServiceType]entity.ErrorMessage

type ServiceType int

const (
	COMMON ServiceType = iota + 1
	HTTP
	SQL
	INTERNAL
	CACHE
)

const (
	// Language Header
	LangEN string = `en`
	LangID string = `id`
)

func init() {
	stacktrace.DefaultFormat = stacktrace.FormatFull
	svcError = map[ServiceType]entity.ErrorMessage{
		COMMON: general.ErrorMessages,
		HTTP:   httperr.ErrorMessages,
		SQL:    sql.ErrorMessages,
		CACHE:  cache.ErrorMessages,
	}
}

// AppError - Application Error Structure
type AppError struct {
	Code       entity.Code `json:"code"`
	Message    string      `json:"message" example:"error"`
	DebugError *string     `json:"debug_error,omitempty" example:"error"`
	sys        error
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *AppError) Error() string {
	return e.sys.Error()
}

// Compile - Get Error Code and HTTP Status
// Common --> Service --> Default
func Compile(service ServiceType, err error, lang string, debugMode bool) (int, AppError) {
	// Developer Debug Error
	var debugErr *string
	if debugMode {
		errStr := err.Error()
		if len(errStr) > 0 {
			debugErr = &errStr
		}
	}

	// Get Error Code
	code := entity.ErrCode(err)

	// Ger Common, HTTP and SQL error message
	types := []ServiceType{COMMON, HTTP, SQL, CACHE}
	for _, st := range types {
		if errMessage, ok := svcError[st][code]; ok {
			msg := errMessage.ID
			if lang == LangEN {
				msg = errMessage.EN
			}

			if errMessage.HasAnnotation {
				args := fmt.Sprintf("%q", err.Error())
				index := strings.Index(args, `\n`)
				if index > 0 {
					args = strings.TrimSpace(args[1:index])
				}
				msg = fmt.Sprintf(msg, args)
			}

			if errMessage.CustomMessage {
				newErr := fmt.Sprintf("%v", err)
				errPart := strings.Split(newErr, "\n")
				newErr = errPart[0]
				msg = fmt.Sprintf(msg, newErr)
			}

			return errMessage.StatusCode, AppError{
				Code:       code,
				Message:    msg,
				sys:        err,
				DebugError: debugErr,
			}
		}
	}

	// Get Service Error
	if errMessages, ok := svcError[service]; ok {
		if errMessage, ok := errMessages[code]; ok {
			msg := errMessage.ID
			if lang == LangEN {
				msg = errMessage.EN
			}

			// error code example
			if errMessage.HasAnnotation {
				args := fmt.Sprintf("%q", err.Error())
				index := strings.Index(args, `\n`)
				if index > 0 {
					args = strings.TrimSpace(args[1:index])
				}
				msg = fmt.Sprintf(msg, args)
			}

			if errMessage.CustomMessage {
				newErr := fmt.Sprintf("%v", err)
				errPart := strings.Split(newErr, "\n")
				newErr = errPart[0]
				msg = fmt.Sprintf(msg, newErr)
			}

			// Humanize Error Msg Fulfillment API
			if code == httperr.CodeHTTPValidatorError {
				if err.Error() != "" {
					msg = strings.Split(err.Error(), "\n ---")[0]
				}
			}

			return errMessage.StatusCode, AppError{
				Code:       code,
				Message:    msg,
				sys:        err,
				DebugError: debugErr,
			}
		}
		return http.StatusInternalServerError, AppError{
			Code:       code,
			Message:    "error message not defined!",
			sys:        err,
			DebugError: debugErr,
		}
	}

	// Set Default Error
	return http.StatusInternalServerError, AppError{
		Code:       code,
		Message:    "service error not defined!",
		sys:        err,
		DebugError: debugErr,
	}
}
