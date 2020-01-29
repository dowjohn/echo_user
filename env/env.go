package env

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func Init() (*Config, error) {
	config := new(Config)

	// Open our jsonFile
	jsonFile, err := os.Open("properties.json")

	if err != nil {
		log.Fatal(err)
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	if err := json.Unmarshal(byteValue, &config); err != nil {
		log.Fatal(err)
	}

	// we dont care if this fails
	if err := jsonFile.Close(); err != nil {
		log.Print(err)
	}

	log.Print("successfully read config")

	return config, err
}

type Config struct {
	PublicKey        string `json:"publicKey"`
	PrivateKey       string `json:"privateKey"`
	AtlasProjectId string `json:"atlasProjectId"`
	AtlasUser string `json:"atlasUser"`
	AtlasPassword string `json:"atlasPassword"`
	AtlasHost string `json:"atlasHost"`
}