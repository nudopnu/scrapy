package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server   Server   `mapstructure:"server"`
	Database Database `mapstructure:"database"`
}

type Server struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	JwtSecret string `mapstructure:"jwt_secret"`
}

type Database struct {
	Host       string `mapstructure:"host"`
	DbName     string `mapstructure:"name"`
	Password   string `mapstructure:"password"`
	DbUserName string `mapstructure:"username"`
}

func LoadConfig() (config *Config, err error) {
	envPrefix := "SCRAPY"
	loadSecrets(envPrefix)
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix(envPrefix)
	viper.AutomaticEnv()

	viper.SetConfigFile("config.toml")
	viper.MergeInConfig()
	viper.SetConfigFile("config-override.toml")
	viper.MergeInConfig()

	config = &Config{}
	if err := viper.Unmarshal(config); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
	return
}

func loadSecrets(envPrefix string) {
	secrets := map[string]string{
		"database.password": "DATABASE_PASSWORD",
		"server.jwt_secret": "SERVER_JWT_SECRET",
	}
	for key, envVar := range secrets {
		// bind secret to environment value
		// this is necessary, because it is not present in config
		// and viper thus not knows about this
		envSecret := fmt.Sprintf("%s_%s", envPrefix, envVar)
		viper.BindEnv(key, envSecret)
		// look for *_FILE environment variables
		envSecretFile := fmt.Sprintf("%s_%s_FILE", envPrefix, envVar)
		secretFile, ok := os.LookupEnv(envSecretFile)
		if !ok {
			continue
		}
		log.Printf("found %s:%s", envSecretFile, secretFile)
		// check if corresponding variable is already set
		_, ok = os.LookupEnv(envSecret)
		if ok {
			continue
		}
		// load secret from file into viper
		value, err := os.ReadFile(secretFile)
		if err != nil {
			log.Fatal(err)
		}
		viper.Set(key, string(value))
	}
}

func (cfg *Config) GetDbUrl() string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", cfg.Database.DbUserName, cfg.Database.Password, cfg.Database.Host, cfg.Database.DbName)
}

func (cfg *Config) GetHostUrl() string {
	return fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
}
