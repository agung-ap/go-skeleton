package common

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	apperr "go-skeleton/pkg/errors"
	x "go-skeleton/pkg/errors/entity"
	"go-skeleton/pkg/errors/general"

	"github.com/gin-gonic/gin"
)

const (
	ContentTypeCSV       = "text/csv"
	ContentTypePlainText = "text/plain"
	ContentTypeJSON      = "application/json"
	ContentTypeXML       = "application/xml"
	ContentTypePDF       = "application/pdf"
	ContentTypeZIP       = "application/zip"
	ContentTypePNG       = "image/png"
	ContentTypeJPEG      = "image/jpeg"
	ContentTypeExcel     = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
)

type Contents struct {
	Description      string
	TransferEncoding string
	Disposition      string
	Types            string
}

// SuccessResp is the structure for all API responses.
type SuccessResp struct {
	Message    string `json:"message"`
	Data       any    `json:"data"`
	Pagination any    `json:"pagination,omitempty"`
}

type ErrorResp struct {
	Error struct {
		Code    x.Code   `json:"code"`
		Message string   `json:"message"`
		Errors  []Errors `json:"errors,omitempty"`
	} `json:"error"`
}

type Errors struct {
	Reason  string `json:"reason"`
	Message string `json:"message"`
}

func ResponseSuccess(c *gin.Context, code int, message string, data any, pagination any) {
	resp := SuccessResp{}
	resp.Message = message
	resp.Data = data

	if pagination != nil {
		resp.Pagination = pagination
	}

	c.JSON(code, resp)
}

func ResponseData(c *gin.Context, code int, content Contents, data []byte) {
	// Set the response headers
	c.Header("Content-Description", content.Description)
	c.Header("Content-Transfer-Encoding", content.TransferEncoding)
	c.Header("Content-Disposition", content.Disposition)
	c.Header("Content-Type", content.Types)

	c.Data(code, content.Types, data)
}

func ResponseError(c *gin.Context, err error, errMessages ...string) {
	debugMode := false
	lang := HeaderLangEN.String()

	if c.GetHeader(HeaderAppDebug.String()) == "true" {
		debugMode = true
	}

	if c.GetHeader(HeaderAppLang.String()) == HeaderLangID.String() {
		lang = HeaderLangID.String()
	}

	// Check if error because context Cancelled or Deadline Exceed
	if errors.Is(x.RootCause(err), context.DeadlineExceeded) {
		err = x.WrapWithCode(err, general.CodeContextDeadlineExceeded, "Error Context Deadline Exceeded")
	}

	if errors.Is(x.RootCause(err), context.Canceled) {
		c.Header(HeaderContentType.String(), "text/plain")
		c.Status(499)
		return
	}

	statusCode, displayError := apperr.Compile(apperr.INTERNAL, err, lang, debugMode)
	statusStr := http.StatusText(statusCode)

	slog.ErrorContext(c, displayError.Error())

	errResp := ErrorResp{
		Error: struct {
			Code    x.Code   `json:"code"`
			Message string   `json:"message"`
			Errors  []Errors `json:"errors,omitempty"`
		}{
			Code:    displayError.Code,
			Message: displayError.Message,
		},
	}

	if len(errMessages) > 0 {
		var errs = make([]Errors, len(errMessages))

		for i, m := range errMessages {
			errs[i].Reason = statusStr
			errs[i].Message = m
		}

		errResp.Error.Errors = errs
	}

	c.JSON(statusCode, errResp)
}
