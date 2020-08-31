package config

type Config struct {
	Server   Server   `yaml:"server"`
	Registry Registry `yaml:"registry"`
	Notify   Notify   `yaml:"notify"`
}
