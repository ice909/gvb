package config

type Mysql struct {
	host     string `yaml:"host"`
	port     int    `yaml:"port"`
	db       string `yaml:"db"`
	user     string `yaml:"user"`
	password string `yaml:"password"`
	logLevel string `yaml:"log-level"`
}
