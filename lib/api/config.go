package api

//-------------------------------------------------------------------------------------------------

var (
	DefaultEnv     = "dev"
	DefaultHost    = "localhost"
	DefaultPort    = 4000
	DefaultAPIPath = "/"
	DefaultWebDir  = "web"
)

//-------------------------------------------------------------------------------------------------

type Config struct {
	Env     string
	Host    string
	Port    int
	APIPath string
	WebDir  string
}

//-------------------------------------------------------------------------------------------------

func NewConfig(env, host string, port int, apiPath, webDir string) *Config {
	return &Config{
		Env:     env,
		Host:    host,
		Port:    port,
		APIPath: apiPath,
		WebDir:  webDir,
	}
}

func DefaultConfig() *Config {
	return NewConfig(
		DefaultEnv,
		DefaultHost,
		DefaultPort,
		DefaultAPIPath,
		DefaultWebDir,
	)
}
