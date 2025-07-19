package sql

import (
	"go-skeleton/pkg/errors/entity"
)

const (
	// Error On SQL
	CodeSQLBuilder = entity.Code(iota + 200)
	CodeSQLRead
	CodeSQLCount
	CodeSQLRowScan
	CodeSQLCreate
	CodeSQLUpdate
	CodeSQLDelete
	CodeSQLUnlink
	CodeSQLTxBegin
	CodeSQLTxRollback
	CodeSQLTxCommit
	CodeSQLPrepareStmt
	CodeSQLRecordMustExist
	CodeSQLCannotRetrieveLastInsertID
	CodeSQLCannotRetrieveAffectedRows
	CodeSQLUniqueConstraint
	CodeSQLRecordDoesNotMatch
	CodeSQLRecordIsExpired
	CodeSQLRecordDoesNotExist
	CodeSQLForeignKeyMissing
	CodeSQLTransactionFailed
	CodeSQLPing
	CodeSQLTruncate
)
