package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string
}

//env-default:"production" means if env variable is not set then default to production

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true" ` //basically we are saying that in the yaml file the field name is env but in struct it is Env and annotations for env variable
	storagePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}
//Mustload dont return error , just do fatal close the application
func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "path to the configuration file")
		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("config path must be provided")
		}
	}


	if _,err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist at path: %s",configPath)
	}

	var cfg Config 

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("failed to read config: %s", err.Error())
	}


	return &cfg
}