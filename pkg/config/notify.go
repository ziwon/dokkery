package config

type Notify struct {
	Slack struct {
		WebHook string `yaml:"webhook"`
		Channel string `yaml:"channel"`
	} `yaml:"slack"`
}
