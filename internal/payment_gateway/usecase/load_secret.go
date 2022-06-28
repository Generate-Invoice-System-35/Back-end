package usecase

import (
	"bytes"

	"io/ioutil"
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Secret struct {
	DB      DBCredential  `yaml:"db"`
	Payment PaymentSecret `yaml:"payment"`
}

type DBCredential struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

type PaymentSecret struct {
	Midtrans APICredential `yaml:"midtrans"`
	Xendit   APICredential `yaml:"xendit"`
}

type APICredential struct {
	ClientID      string `yaml:"clientId"`
	ClientKey     string `yaml:"clientKey"`
	SecretKey     string `yaml:"secretKey"`
	CallbackToken string `yaml:"callbackToken"`
}

func LoadSecret(path string) (*Secret, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return LoadSecretFromBytes(data)
}

func LoadSecretFromBytes(data []byte) (*Secret, error) {
	fang := viper.New()
	fang.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	fang.AutomaticEnv()
	fang.SetEnvPrefix("GOPAYMENT")
	fang.SetConfigType("yaml")

	if err := fang.ReadConfig(bytes.NewBuffer(data)); err != nil {
		return nil, err
	}

	var creds Secret
	err := fang.Unmarshal(&creds)
	if err != nil {
		log.Fatalf("Error loading creds: %v", err)
	}

	return &creds, nil
}
