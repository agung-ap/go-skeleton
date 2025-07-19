package http

import (
	errors "go-skeleton/pkg/errors/entity"
)

var ErrorMessages = errors.ErrorMessage{
	CodeHTTPClientMarshal:         errors.ErrMsgISE,
	CodeHTTPClientUnmarshal:       errors.ErrMsgISE,
	CodeHTTPClientErrorOnRequest:  errors.ErrMsgISE,
	CodeHTTPClientErrorOnReadBody: errors.ErrMsgISE,
	CodeHTTPInternalServerError:   errors.ErrMsgISE,
	CodeHTTPNotFound:              errors.ErrMsgNotFound,
	CodeHTTPBadRequest:            errors.ErrMsgBadRequest,
	CodeHTTPBadRequestCustom:      errors.ErrMsgBadRequestCustom,
	CodeHTTPUnauthorized:          errors.ErrMsgUnauthorized,
	CodeHTTPUnmarshal:             errors.ErrMsgBadRequest,
	CodeHTTPUnmarshalCustom:       errors.ErrMsgBadRequestCustom,
	CodeHTTPMarshal:               errors.ErrMsgISE,
	CodeHTTPConflict:              errors.ErrMsgConflict,
	CodeHTTPForbidden:             errors.ErrMsgForbidden,
	CodeHTTPServiceUnavailable:    errors.ErrMsgServiceUnavailable,
	CodeHTTPVersionConstraint:     errors.ErrMsgVersionConstraint,
	CodeHTTPParamDecode:           errors.ErrMsgBadRequest,
	CodeHTTPErrorOnReadBody:       errors.ErrMsgISE,
	CodeHTTPTooManyRequest:        errors.ErrMsgTooManyRequest,
}
