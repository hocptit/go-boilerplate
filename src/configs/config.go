package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port                    string `mapstructure:"PORT"`
	DBUrl                   string `mapstructure:"DB_URL"`
	Env                     string `mapstructure:"ENV"`
	AppName                 string `mapstructure:"APP_NAME"`
	AppIsWriteLog           string `mapstructure:"APP_IS_WRITE_LOG"`
	AppIsReturnDetailErrors string `mapstructure:"APP_IS_RETURN_DETAIL_ERRORS"`
	DBIsWriteLog            string `mapstructure:"DB_IS_WRITE_LOG"`
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
