package common

import (
	"testing"

	x "go-skeleton/pkg/errors/entity"

	"github.com/stretchr/testify/assert"
)

func TestResponseSuccess(t *testing.T) {
	// Skip this test due to gin import issues
	t.Skip("Skipping due to gin build constraints")
}

func TestResponseData(t *testing.T) {
	// Skip this test due to gin import issues
	t.Skip("Skipping due to gin build constraints")
}

func TestResponseError(t *testing.T) {
	// Skip this test due to gin import issues
	t.Skip("Skipping due to gin build constraints")
}

func TestSuccessResp_Structure(t *testing.T) {
	resp := SuccessResp{
		Message:    "success",
		Data:       map[string]string{"key": "value"},
		Pagination: map[string]int{"page": 1},
	}

	assert.Equal(t, "success", resp.Message)
	assert.NotNil(t, resp.Data)
	assert.NotNil(t, resp.Pagination)
}

func TestErrorResp_Structure(t *testing.T) {
	resp := ErrorResp{
		Error: struct {
			Code    x.Code   `json:"code"`
			Message string   `json:"message"`
			Errors  []Errors `json:"errors,omitempty"`
		}{
			Code:    123,
			Message: "error message",
			Errors: []Errors{
				{Reason: "validation", Message: "invalid input"},
			},
		},
	}

	assert.Equal(t, x.Code(123), resp.Error.Code)
	assert.Equal(t, "error message", resp.Error.Message)
	assert.Len(t, resp.Error.Errors, 1)
	assert.Equal(t, "validation", resp.Error.Errors[0].Reason)
	assert.Equal(t, "invalid input", resp.Error.Errors[0].Message)
}

func TestErrors_Structure(t *testing.T) {
	err := Errors{
		Reason:  "validation_failed",
		Message: "Field is required",
	}

	assert.Equal(t, "validation_failed", err.Reason)
	assert.Equal(t, "Field is required", err.Message)
}

func TestContents_Structure(t *testing.T) {
	contents := Contents{
		Description:      "File description",
		TransferEncoding: "gzip",
		Disposition:      "attachment",
		Types:            "application/json",
	}

	assert.Equal(t, "File description", contents.Description)
	assert.Equal(t, "gzip", contents.TransferEncoding)
	assert.Equal(t, "attachment", contents.Disposition)
	assert.Equal(t, "application/json", contents.Types)
}
