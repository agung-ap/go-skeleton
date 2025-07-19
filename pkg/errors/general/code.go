package general

import (
	"go-skeleton/pkg/errors/entity"
)

const (
	// Error Common
	CodeValueInvalid = entity.Code(iota + 1000)
	CodeContextDeadlineExceeded
	CodeContextCanceled

	// CodeFileOperationError File Operation Errors
	CodeFileOperationError
	CodeFileCreateError
	CodeFileOpenError
	CodeFileReadError
	CodeFileWriteError
	CodeFileCloseError
	CodeFileRemoveError
	CodeFileStatError
	CodeFilePermissionError

	// CodeCmdExecError Command Execution Errors
	CodeCmdExecError
	CodeCmdStartError
	CodeCmdRunError
	CodeCmdWaitError
	CodeCmdPipeError
	CodeCmdTimeoutError
)
