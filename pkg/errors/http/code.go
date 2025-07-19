package http

import "go-skeleton/pkg/errors/entity"

const (
	// Code HTTP Handler
	CodeHTTPBadRequest = entity.Code(iota + 800)
	CodeHTTPBadRequestCustom
	CodeHTTPNotFound
	CodeHTTPUnauthorized
	CodeHTTPInternalServerError
	CodeHTTPUnmarshal
	CodeHTTPUnmarshalCustom
	CodeHTTPMarshal
	CodeHTTPConflict
	CodeHTTPForbidden
	CodeHTTPTooManyRequest
	CodeHTTPValidatorError
	CodeHTTPServiceUnavailable
	CodeHTTPVersionConstraint
	CodeHTTPParamDecode
	CodeHTTPErrorOnReadBody
)

const (
	// Error on HTTP Client
	CodeHTTPClientMarshal = entity.Code(iota + 500)
	CodeHTTPClientUnmarshal
	CodeHTTPClientErrorOnRequest
	CodeHTTPClientErrorOnReadBody
)
