package sql

import (
	errors "go-skeleton/pkg/errors/entity"
)

var ErrorMessages = errors.ErrorMessage{
	CodeSQLBuilder:                    errors.ErrMsgISE,
	CodeSQLRead:                       errors.ErrMsgISE,
	CodeSQLCount:                      errors.ErrMsgISE,
	CodeSQLRowScan:                    errors.ErrMsgISE,
	CodeSQLCreate:                     errors.ErrMsgISE,
	CodeSQLUpdate:                     errors.ErrMsgISE,
	CodeSQLDelete:                     errors.ErrMsgISE,
	CodeSQLUnlink:                     errors.ErrMsgISE,
	CodeSQLTxBegin:                    errors.ErrMsgISE,
	CodeSQLTxRollback:                 errors.ErrMsgISE,
	CodeSQLTxCommit:                   errors.ErrMsgISE,
	CodeSQLPrepareStmt:                errors.ErrMsgISE,
	CodeSQLRecordMustExist:            errors.ErrMsgNotFound,
	CodeSQLCannotRetrieveLastInsertID: errors.ErrMsgISE,
	CodeSQLCannotRetrieveAffectedRows: errors.ErrMsgISE,
	CodeSQLUniqueConstraint:           errors.ErrMsgUniqueConst,
	CodeSQLRecordDoesNotMatch:         errors.ErrMsgBadRequest,
	CodeSQLRecordIsExpired:            errors.ErrMsgBadRequest,
	CodeSQLPing:                       errors.ErrMsgBadRequestCustom,
	CodeSQLRecordDoesNotExist:         errors.ErrMsgNotFound,
	CodeSQLForeignKeyMissing:          errors.ErrMsgISE,
	CodeSQLTransactionFailed:          errors.ErrMsgISE,
	CodeSQLTruncate:                   errors.ErrMsgISE,
}
