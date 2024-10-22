package config

type System struct {
	host string `yaml:"host"`
	port int    `yaml:"port"`
	env  string `yaml:"env"`
}
