package goconfig

import "github.com/Unknwon/goconfig"

type Config struct {
	cfg *goconfig.ConfigFile
}

func load(fileName string) (*Config, error) {
	if cfg, err := goconfig.LoadConfigFile(fileName); err != nil {
		return nil, err
	} else {
		return &Config{
			cfg: cfg,
		}, nil
	}
}

func (c *Config) getFromConfig(key string) (string, error) {
	return c.cfg.GetValue("", key)
}
