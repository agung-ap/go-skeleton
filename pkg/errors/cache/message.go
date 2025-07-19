package cache

import (
	errors "go-skeleton/pkg/errors/entity"
)

var ErrorMessages = errors.ErrorMessage{
	CodeCacheRead:            errors.ErrMsgISE,
	CodeCacheCount:           errors.ErrMsgISE,
	CodeCacheCreate:          errors.ErrMsgISE,
	CodeCacheUpdate:          errors.ErrMsgISE,
	CodeCacheDelete:          errors.ErrMsgISE,
	CodeCacheMustExist:       errors.ErrMsgNotFound,
	CodeCacheDoesNotMatch:    errors.ErrMsgBadRequest,
	CodeCacheIsExpired:       errors.ErrMsgBadRequest,
	CodeCacheDoesNotExist:    errors.ErrMsgNotFound,
	CodeCacheDecode:          errors.ErrMsgISE,
	CodeCacheMarshal:         errors.ErrMsgISE,
	CodeCacheUnmarshal:       errors.ErrMsgISE,
	CodeCacheDeleteSimpleKey: errors.ErrMsgISE,
}
