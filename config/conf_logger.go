package config

type Logger struct {
	level        string `yaml:"level"`
	prefix       string `yaml:"prefix"`
	director     string `yaml:"director"`
	showLine     bool   `yaml:"show-line"`
	logInConsole bool   `yaml:"log-in-console"`
}
