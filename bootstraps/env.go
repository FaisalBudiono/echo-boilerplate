package bootstraps

import "github.com/spf13/viper"

func BindEnv() {
	viper.AutomaticEnv()
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
}
