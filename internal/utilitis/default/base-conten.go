package _default

import "github.com/labstack/echo/v4"

type BaseContent struct {
	Code        int
	Status      string
	Description string
	Data        interface{}
}

func BaseHandlerReturn(ctx echo.Context, content *BaseContent) error {

}
