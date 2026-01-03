package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr  string `yaml:"address" env-required:"true"`
}

//env-default:"production" means if env variable is not set then default to production

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true" ` //basically we are saying that in the yaml file the field name is env but in struct it is Env and annotations for env variable
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}
//Mustload dont return error , just do fatal close the application
func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")  // basically we are checking if config path is set in env variable for the command we write in terminal

	if configPath == "" {   //checking if env is empty then go inside condition and parse command line arguments
		flags := flag.String("config", "", "path to the configuration file")
		flag.Parse()  //parsing the command line arguments

		configPath = *flags  //dereferencing the pointer to get the actual value

		if configPath == "" {
			log.Fatal("config path must be provided")  //if config path is still empty then we log fatal and exit
		}
	}


	if _,err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist at path: %s",configPath)  //checking if the file exists at the given path
	}

	var cfg Config   //creating an instance of config struct

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("failed to read config: %s", err.Error()) //reading the config file and populating the struct
	}


	return &cfg  //returning the pointer to the config struct
}