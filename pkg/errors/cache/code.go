package cache

import (
	"go-skeleton/pkg/errors/entity"
)

const (
	CodeCacheRead = entity.Code(iota + 100)
	CodeCacheCount
	CodeCacheCreate
	CodeCacheUpdate
	CodeCacheDelete
	CodeCacheMustExist
	CodeCacheDoesNotMatch
	CodeCacheIsExpired
	CodeCacheDoesNotExist
	CodeCacheDecode
	CodeCacheMarshal
	CodeCacheUnmarshal
	CodeCacheDeleteSimpleKey
)
