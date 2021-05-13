package main

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ProjectIdFirestore            string  `envconfig:"PROJECT_ID_FIRESTORE" required:"true"`
	FirebaseServiceAccountKeyJson string  `envconfig:"FIREBASE_SERVICE_ACCOUNT_KEY_JSON" required:"true"`
}


func NewConfig() (*Config, error) {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}