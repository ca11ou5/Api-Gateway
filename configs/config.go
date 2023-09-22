package configs

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Port       string `mapstructure:"PORT"`
	AuthSvcUrl string `mapstructure:"AUTH_SVC_URL"`
}

func InitConfig() (c Config) {
	viper.SetConfigFile("configs/envs/dev.env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("error when reading config")
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		log.Fatal("error when unmarshalling config")
	}

	return
}
