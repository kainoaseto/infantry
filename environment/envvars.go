package environment

import (
	"github.com/kelseyhightower/envconfig"
)

// TODO: Currently this just processes entire configs via structures. This sucks.
//       It needs to be able to do independent dynamic updating of environment variables
//       set in the kernel and via config file.

//       Since no one does this well anymore (RIP Viper) we will create our own here and possible
//       export this package out to another repository in the future. We need to be able to read
//       INI files since that is how docker manages .env files that we want to pass it. Here is a good
//       package for this: https://github.com/go-ini/ini

//       For dynamicness some of the work will need to be done by us, for the rest lets use Consul
//       and this sweet thread safe thing called MVGA: https://github.com/rbastic/mvga
//       That code is a mess so we will just credit the guy and rewrite it all ourself taking excerpts
//       from it.

type VariableManager struct {
	config struct{}
}

func (*VariableManager) Get(key stirng) {

}

func (*VariableManager) Process(config struct{}) {
	envconfig.Process("")
}
