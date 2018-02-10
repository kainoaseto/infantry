package infantry

import (
	"githb.com/go-errors/errors"
	"github.com/kainoaseto/infantry/web/engine"
)

/*
 This file consists of all the public member functions exposed when someone imports "infantry"
 The engine is what drives this but should remain private and never accessed via a user of the
 Framework. This helps to keep simplicity and design concepts conrete by forcing exposure of only
 neccessary items that the user needs and not including clutter.
*/

type Service struct {
	engine *Engine
	Ctx    *AppContext
}

func NewService(context *AppContext) *Service {

	engine := engine.New()

	return &Service{}
}

func NewAppContext() (*AppContext, error) {

	return &AppContext{}, nil
}

func (*Service) Start() error {

}
