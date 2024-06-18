package utils

import "github.com/spf13/viper"

type Config struct {
	DBHostname string `mapstructure:"DB_HOSTNAME"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
}

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()
	err := viper.ReadInConfig()

	if err != nil {
		return nil, err
	}

	var config *Config

	if err = viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return config, nil
}
