package util

import "github.com/spf13/viper"

type Config struct{
	DB_SOURCE string `mapstructure:"DB_SOURCE"`
	DB_DRIVE string `mapstructure:"DB_DRIVER"`
	SERVER_ADDRESS string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(pathname string)(config Config,err error){
	viper.AddConfigPath(pathname)
	viper.SetConfigType("env")
	viper.SetConfigName("app")
	viper.AutomaticEnv()

	err=viper.ReadInConfig()
	if err!=nil{
		return
	}

	err=viper.Unmarshal(&config)
	return

}