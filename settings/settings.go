package settings

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

// Setup read the settings from the settings file and return an entities.Settings
func Setup() (*Settings, error) {
	file, err := os.Open("settings.yml")
	if err != nil {
		log.Println("[Setup] Error Open", err)
		return nil, err
	}

	var settings Settings
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&settings)
	if err != nil {
		log.Println("[Setup] Error Decode", err)
		return nil, err
	}

	err = file.Close()
	if err != nil {
		log.Println("[Setup] Close", err)
		return nil, err
	}

	return &settings, nil
}
