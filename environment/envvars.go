package environment

// TODO: Soooo after some research we want to use Viper and possibly contribute to it or Fork it and do a vast revision
// with pullin gin PRs from Viper that desperaretly need to get merged. At this point we may want to change EnvVarsManager
// to ConfigManager or create a configuration package outside of environment and then setup Viper's environment management
// to monitor the .env file and handle setting up Environemt variables and all that. Actually the more I write about it the
// more I like the second option. Lets start doing that unless there is opposition

type VariableManager struct {
}
