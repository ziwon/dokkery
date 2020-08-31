package config

type Notify struct {
	Slack struct {
		WebHook string `yaml:"webhook"`
		Channel string `yaml:"channel"`
		Message struct {
			Success struct {
				Head string `yaml:"head"`
			} `yaml:"success"`
			Fail struct {
				Head string `yaml:"head"`
			} `yaml:"fail"`
		} `yaml:"message"`
	} `yaml:"slack"`
}
