package engine

import (
	"strings"

	"github.com/go-errors/errors"

	"github.com/kainoaseto/infantry/clients"
	"github.com/kainoaseto/infantry/environment"
	"github.com/kainoaseto/infantry/logger"
)

/*
 * Anything marked with "RO" will have a default that is grabbed from an environment variable and the setting is
 * exists simply for accessiblity
 */
type (
	ServerSettings struct {
		Port        string // Port that the service will be accesible on
		Environment string // RO: The name of the environment that this will will live in
		Cluster     string // RO: The name of the service cluster that this will live in
		App         string // RO: The name of the service
	}

	Kernel struct {
		Server  ServerSettings // Server Settings defined above to be set before server creation
		Env     EnvManager     // Environment manager to handle all things with the env variables
		Clients ClientManager  // Client manager for all external communication with known services
		Logger  LoggingManager // LoggingManager that handles all things logging related
	}
)

func (k *Kernel) NewKernel() (*Kernel, error) {

}
