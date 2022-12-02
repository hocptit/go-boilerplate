package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port  string `mapstructure:"PORT"`
	DbUrl string `mapstructure:"DB_URL"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.SetConfigFile(path)
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
