package configurator

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func init() {
	mode := os.Getenv("MODE")
	if mode == "" || (mode != "development" && mode != "production") {
		mode = "development"
	}
	filename := fmt.Sprintf("config.%s", mode)
	viper.SetConfigName(filename)
	viper.AddConfigPath("./config/")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func ReadConfig(target interface{}) error {
	if err := viper.Unmarshal(&target); err != nil {
		return err
	}
	return nil
}

func ReadConfigByKey(key string, target interface{}) error {
	if err := viper.UnmarshalKey(key, &target); err != nil {
		return err
	}
	return nil
}
