package bootstraps

import (
	"time"

	"github.com/spf13/viper"
)

func getAddr() string {
	return ":" + viper.GetString("PORT")
}

func getAppName() string {
	return viper.GetString("APP_NAME")
}

func getHttpTimeout() time.Duration {
	return time.Second * time.Duration(viper.GetInt("HTTP_TIMEOUT"))
}
