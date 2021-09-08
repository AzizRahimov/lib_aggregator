package conf

import "github.com/spf13/viper"

func InitConfig() error {



	viper.AddConfigPath("conf")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
