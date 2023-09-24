package config

import "github.com/spf13/viper"

type Config struct {
	Port       string `mapstructure:"PORT"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
}

var envs = []string{"DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME"}

func LoadConfig() (config Config, err error) {

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")

	viper.ReadInConfig()

	for _, env := range envs {

		if err = viper.BindEnv(env); err != nil {
			return Config{}, err
		}
	}

	if err = viper.Unmarshal(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}
