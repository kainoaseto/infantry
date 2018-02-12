package recon

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Loads a config into memory and into the environment variables
// Use's godotenv to load the .env into the environment variables and
// envconfig for marshalling the config types properly.
// Params: configPath, envConfig spec address, option appName for prefix on envs
func LoadConfig(configPath string, configSpec interface{}, appName ...string) error {

	err := envconfig.Process(appName[0], configSpec)
	if err != nil {
		return err
	}

	err = godotenv.Load(configPath)
	if err != nil {
		return err
	}

	return nil
}
