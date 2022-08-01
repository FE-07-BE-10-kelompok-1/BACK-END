package config

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Driver                string
	Name                  string
	Address               string
	Port                  int
	Username              string
	Password              string
	AWS_REGION            string
	AWS_ACCESS_KEY_ID     string
	AWS_SECRET_ACCESS_KEY string
	BUCKET_NAME           string
	MIDTRANS_SERVER_KEY   string
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig

	if _, exist := os.LookupEnv("SECRET"); !exist {
		if err := godotenv.Load("setup.env"); err != nil {
			log.Fatal(err)
		}
	}

	SECRET = os.Getenv("SECRET")
	cnv, err := strconv.Atoi(os.Getenv("SERVERPORT"))
	if err != nil {
		log.Fatal("Cannot parse port variable")
		return nil
	}
	SERVERPORT = int16(cnv)
	defaultConfig.Name = os.Getenv("DB_NAME")
	defaultConfig.Username = os.Getenv("DB_USERNAME")
	defaultConfig.Password = os.Getenv("DB_PASSWORD")
	defaultConfig.Address = os.Getenv("DB_ADDRESS")
	cnv, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("Cannot parse DB Port variable")
		return nil
	}
	defaultConfig.Port = cnv
	defaultConfig.AWS_REGION = os.Getenv("AWS_REGION")
	defaultConfig.AWS_ACCESS_KEY_ID = os.Getenv("AWS_ACCESS_KEY_ID")
	defaultConfig.AWS_SECRET_ACCESS_KEY = os.Getenv("AWS_SECRET_ACCESS_KEY")
	defaultConfig.BUCKET_NAME = os.Getenv("BUCKET_NAME")
	defaultConfig.MIDTRANS_SERVER_KEY = os.Getenv("MIDTRANS_SERVER_KEY")
	return &defaultConfig
}
