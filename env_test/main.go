package main

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

var (
	config *Config
)

func main() {

	//os.Setenv("AGE", "1")

	cfg := GetConfig()

	fmt.Println(cfg.Age)

	fmt.Println(cfg)

}

type Config struct {
	Age int `envconfig:"AGE" required:"true"`
}


func GetConfig() *Config {
	if config == nil {
		config = new(Config)
		err := envconfig.Process("", config)
		if err != nil {
			zap.L().Panic("load env fail", zap.String("error",err.Error()))
		}
	}
	return config
}
