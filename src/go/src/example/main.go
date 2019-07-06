package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()
	fmt.Printf("%s", viper.GetString("HOME"))
}
