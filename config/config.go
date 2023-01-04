package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

var config Config
var doOnce sync.Once

type Config struct {
	HTTPPort         int    `mapstructure:"HTTP_PORT"`
	InClusterAccess  bool   `mapstructure:"IN_CLUSTER_ACCESS"`
	ClusterNamespace string `mapstructure:"CLUSTER_NAMESPACE"`
}

func Get() Config {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("cannot read .env file: %v", err)
	}

	doOnce.Do(func() {
		err := viper.Unmarshal(&config)
		if err != nil {
			log.Fatalln("cannot unmarshal config")
		}
	})

	return config
}
