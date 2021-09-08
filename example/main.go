package main

import (
	"fmt"

	"github.com/bitmark-inc/config-loader"
	"github.com/spf13/viper"
)

func main() {
	config.LoadConfig("MY_CONF")
	purpose := viper.GetString("purpose")
	fmt.Println("Purpose:", purpose)
}
