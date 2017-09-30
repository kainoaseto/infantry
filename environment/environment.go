package environment

import (
	"github.com/go-errors/errors"
)

/*
 * Package: Environment
 * Purpose: To allow communication with the underlying host environment that the go binary runs on
 */

type EnvironmentManager struct {
	Vars VariableManager // Environment Variables Manager to allow Sets and Gets for Env vars
}

func (em *EnvironmentManager) NewEnvironmentManager() (*EnvironmentManager, error) {

}
