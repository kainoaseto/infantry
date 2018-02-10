package engine

import (
	"github.com/labstack/echo"
)

type Engine struct {
	echoBase *Echo
}

func NewEngine() (*Engine, error) {
	e := echo.New()

	return &Engine{}, nil
}

func (*Engine) Use() {

}

func (*Engine) Start() {

}
