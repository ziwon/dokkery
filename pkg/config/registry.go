package config

type Registry struct {
	Domain string `yaml:"domain"`

	OnPush struct {
		Services []struct {
			Name  string   `yaml:"name"`
			Image string   `yaml:"image"`
			Pre   []string `yaml:"pre"`
			Post  []string `yaml:"post"`
		} `yaml:"services"`
	} `yaml:"onpush"`

	OnPull struct {
		Services []struct {
			Name  string   `yaml:"name"`
			Image string   `yaml:"image"`
			Pre   []string `yaml:"pre"`
			Post  []string `yaml:"post"`
		} `yaml:"services"`
	} `yaml:"onpull"`

	OnDelete struct {
		Services []struct {
			Name  string   `yaml:"name"`
			Image string   `yaml:"image"`
			Pre   []string `yaml:"pre"`
			Post  []string `yaml:"post"`
		} `yaml:"services"`
	} `yaml:"onpull"`
}
