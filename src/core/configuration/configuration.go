package configuration

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"main/auth"
	"os"
)

// Builds Config Struct because im lazy
//go:generate /Users/carterloyd/GolandProjects/OpenSOAR/scripts/structConstructor -t yaml -n Configuration -p configuration -f /Users/carterloyd/GolandProjects/OpenSOAR/bin/config.yaml

func Load(configPath string) (Configuration, error) {

	/// Read in the YAML file
	data, err := os.ReadFile(configPath)
	if err != nil {
		return Configuration{}, fmt.Errorf("unable to open file: %v", err)
	}
	// Unmarshal the YAML data into your struct
	var config Configuration
	if err := yaml.Unmarshal(data, &config); err != nil {
		return Configuration{}, fmt.Errorf("Unable to Unmarshal YAML: %v", err)
	}

	// Generate JWT Token Key
	err = auth.GenerateKey(32)

	return config, nil
}
