package config

import "github.com/spf13/viper"

const (
	defaultServerHost = "127.0.0.1"
	defaultServerPort = 8082
	defaultBaseDir    = "/var/kotf"
)

func Init() {
	viper.SetConfigName("app")
	viper.SetConfigType("yml")
	viper.SetDefault("server", server{
		host: defaultServerHost,
		port: defaultServerPort,
	})
	viper.SetDefault("base", defaultBaseDir)
	viper.AddConfigPath("/etc/kotf")
	viper.AddConfigPath("./")
	_ = viper.ReadInConfig()
}

type server struct {
	host string
	port int
}
