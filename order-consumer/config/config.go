package config

import "os"

type Config struct {
	Application string
	Environment string
	MachineName string
	ApiPort     string
	RabbitUrl   string
	MongoUrl    string
}

type Params struct {
	DevMode bool
}

var conf Config

func Get() Config {
	return conf
}

func GetConfig() {
	conf = Config{
		Application: os.Getenv("APPLICATION"),
		Environment: os.Getenv("ENVIRONMENT"),
		MachineName: os.Getenv("MACHINE_NAME"),
		ApiPort:     os.Getenv("API_PORT"),
		RabbitUrl:   os.Getenv("RABBIT_URL"),
		MongoUrl:    os.Getenv("MONGO_URL"),
	}
}

func SetConfig(params *Params) {
	if params.DevMode {
		os.Setenv("APPLICATION", "OrderConsumer")
		os.Setenv("ENVIRONMENT", "Development")
		os.Setenv("MACHINE_NAME", getHostName())
		os.Setenv("API_PORT", "9191")
		os.Setenv("RABBIT_URL", "amqp://guest:guest@localhost:5672/")
		os.Setenv("MONGO_URL", "localhost:27017")
	}
	GetConfig()
}

func getHostName() string {
	machineName, err := os.Hostname()
	if err != nil {
		return ""
	}
	return machineName
}
