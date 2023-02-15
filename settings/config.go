package settings

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/spf13/viper"
)

type (
	Config struct {
		AppEnv                 string `mapstructure:"APP_ENV"`
		ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
		DBHost                 string `mapstructure:"POSTGRES_HOST"`
		DBPort                 string `mapstructure:"POSTGRES_PORT"`
		DBUser                 string `mapstructure:"POSTGRES_USER"`
		DBPass                 string `mapstructure:"POSTGRES_PASSWORD"`
		DBName                 string `mapstructure:"POSTGRES_DB"`
		AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
		RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
		AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
		RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
	}
)

var (
	AppConfig   Config
	onceEnvLoad sync.Once
)

func ConvertToInt(stringNumber string) int {

	if stringNumber == "" {
		return 2
	}

	i, err := strconv.Atoi(stringNumber)
	if err != nil {
		log.Fatal(err)
		return 2
	}

	return i
}

func AskForEnv() {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if AppConfig.AppEnv == "development" || AppConfig.AppEnv == "dev" {
		log.Println("The App is running in development env")
	}
}

func SetConfig() {

	onceEnvLoad.Do(func() {
		if os.Getenv("IS_CONTAINER") != "TRUE" {
			AskForEnv()
		} else {
			AppConfig.AppEnv = os.Getenv("APP_ENV")
			AppConfig.ServerAddress = os.Getenv("SERVER_ADDRESS")
			AppConfig.DBHost = os.Getenv("POSTGRES_HOST")
			AppConfig.DBPort = os.Getenv("POSTGRES_PORT")
			AppConfig.DBUser = os.Getenv("POSTGRES_USER")
			AppConfig.DBPass = os.Getenv("POSTGRES_PASSWORD")
			AppConfig.DBName = os.Getenv("POSTGRES_DB")
			AppConfig.AccessTokenExpiryHour = ConvertToInt(os.Getenv("ACCESS_TOKEN_EXPIRY_HOUR"))
			AppConfig.RefreshTokenExpiryHour = ConvertToInt(os.Getenv("REFRESH_TOKEN_EXPIRY_HOUR"))
			AppConfig.AccessTokenSecret = os.Getenv("ACCESS_TOKEN_SECRET")
			AppConfig.RefreshTokenSecret = os.Getenv("REFRESH_TOKEN_SECRET")
		}

	})

}