package general

import (
	"net/http"

	errors "go-skeleton/pkg/errors/entity"
)

var ErrorMessages = errors.ErrorMessage{
	CodeValueInvalid:            errors.ErrMsgBadRequest,
	CodeContextDeadlineExceeded: errors.ErrMsgContextTimeout,
	CodeContextCanceled:         errors.ErrMsgContextCancelled,

	// File Operation Errors
	CodeFileOperationError: {
		StatusCode: http.StatusInternalServerError,
		EN:         `File operation failed. Please try again or contact support if the problem persists.`,
		ID:         `Operasi file gagal. Silakan coba lagi atau hubungi dukungan jika masalah berlanjut.`,
	},
	CodeFileCreateError: {
		StatusCode: http.StatusInternalServerError,
		EN:         `Unable to create file. Please check permissions and try again.`,
		ID:         `Tidak dapat membuat file. Silakan periksa izin dan coba lagi.`,
	},
	CodeFileOpenError: {
		StatusCode: http.StatusInternalServerError,
		EN:         `Unable to open file. Please ensure the file exists and is accessible.`,
		ID:         `Tidak dapat membuka file. Silakan pastikan file ada dan dapat diakses.`,
	},
	CodeFileReadError: {
		StatusCode: http.StatusInternalServerError,
		EN:         `Unable to read file. Please check file permissions and try again.`,
		ID:         `Tidak dapat membaca file. Silakan periksa izin file dan coba lagi.`,
	},
	CodeFileWriteError: {
		StatusCode: http.StatusInternalServerError,
		EN:         `Unable to write to file. Please check disk space and permissions.`,
		ID:         `Tidak dapat menulis ke file. Silakan periksa ruang disk dan izin.`,
	},
	CodeFileRemoveError: {
		StatusCode: http.StatusInternalServerError,
		EN:         `Unable to remove file. Please check permissions and try again.`,
		ID:         `Tidak dapat menghapus file. Silakan periksa izin dan coba lagi.`,
	},
	CodeFileStatError: {
		StatusCode: http.StatusInternalServerError,
		EN:         `Unable to get file information. Please check file permissions and try again.`,
		ID:         `Tidak dapat mendapatkan informasi file. Silakan periksa izin file dan coba lagi.`,
	},
	CodeFilePermissionError: {
		StatusCode: http.StatusInternalServerError,
		EN:         `File permission error. Please check file permissions and try again.`,
		ID:         `Kesalahan izin file. Silakan periksa izin file dan coba lagi.`,
	},

	// Command Execution Errors
	CodeCmdExecError: {
		StatusCode: http.StatusInternalServerError,
		EN:         `Command execution failed. Please try again or contact support.`,
		ID:         `Eksekusi perintah gagal. Silakan coba lagi atau hubungi dukungan.`,
	},
	CodeCmdStartError: {
		StatusCode: http.StatusInternalServerError,
		EN:         `Unable to start command. Please check system resources and try again.`,
		ID:         `Tidak dapat memulai perintah. Silakan periksa sumber daya sistem dan coba lagi.`,
	},
	CodeCmdRunError: {
		StatusCode: http.StatusInternalServerError,
		EN:         `Unable to run command. Please check command parameters and try again.`,
		ID:         `Tidak dapat menjalankan perintah. Silakan periksa parameter perintah dan coba lagi.`,
	},
	CodeCmdWaitError: {
		StatusCode: http.StatusInternalServerError,
		EN:         `Command execution was interrupted. Please try again.`,
		ID:         `Eksekusi perintah terganggu. Silakan coba lagi.`,
	},
	CodeCmdPipeError: {
		StatusCode: http.StatusInternalServerError,
		EN:         `Unable to create command pipe. Please try again or contact support.`,
		ID:         `Tidak dapat membuat pipa perintah. Silakan coba lagi atau hubungi dukungan.`,
	},
	CodeCmdTimeoutError: {
		StatusCode: http.StatusInternalServerError,
		EN:         `Command execution timed out. Please try again with a simpler operation.`,
		ID:         `Waktu eksekusi perintah habis. Silakan coba lagi dengan operasi yang lebih sederhana.`,
	},
}
