package QueryData

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Secret struct {
	AccessTokenQu   string `json:"accessTokenQu"`
	AccessTokenIsbn string `json:"accessTokenIsbn"`
}

func LoadSecret() {
	fi, err := os.Stat("secret.json")
	if err != nil {
		log.Printf("Loss secret file")
		return
	}
	file, err := os.Open(fi.Name())
	defer file.Close()

	configBytes, err := io.ReadAll(file)
	err = json.Unmarshal(configBytes, &secret)
	if err != nil {
		log.Printf(err.Error())
		return
	}
}
