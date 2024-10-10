package _default

import "net/http"

type BaseContent struct {
	Code        int
	Status      string
	Description string
	Data        interface{}
}

func InternalError(err error) *BaseContent {
	return &BaseContent{
		Code:        http.StatusInternalServerError,
		Status:      "failed",
		Description: err.Error(),
		Data:        nil,
	}
}
