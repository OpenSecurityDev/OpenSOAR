package configuration

// CODE GENERATED WITH GO GENERATE
// DO NOT EDIT BY HAND

type Configuration struct {
	Authentication struct {
		JWTSecret   string `yaml:"JWT_secret"`
		LoginMethod string `yaml:"login_method"`
	} `yaml:"authentication"`
	Configuration struct {
		PluginDirectory string `yaml:"plugin_directory"`
		Server          struct {
			Address string `yaml:"address"`
			Port    int    `yaml:"port"`
		} `yaml:"server"`
	} `yaml:"configuration"`
	Logging struct {
		Format string `yaml:"format"`
		Level  string `yaml:"level"`
		Output struct {
			Destination string `yaml:"destination"`
		} `yaml:"output"`
	} `yaml:"logging"`
}
