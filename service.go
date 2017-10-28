package infantry

import (
	"github.com/kainoaseto/infantry/web/engine"
	"error"
)

/*
 This file consists of all the public member functions exposed when someone imports "infantry"
 The engine is what drives this but should remain private and never accessed via a user of the
 Framework. This helps to keep simplicity and design concepts conrete by forcing exposure of only
 neccessary items that the user needs and not including clutter.
*/

type Service struct {
}

func NewService(kernel *Kernel) *Service {

	return &Service{}
}

func NewKernel() *Kernel, error {
	return &Kernel{}, nil
}
