package swagger

import (
	"github.com/linshenqi/sptty"
	"github.com/linshenqi/sptty.swagger/src/base"
)

type Config struct {
	sptty.BaseConfig

	Enable bool   `yaml:"enable"`
	Url    string `yaml:"url"`
	Auto   bool   `yaml:"auto"`
}

func (s *Config) ConfigName() string {
	return base.ServiceSwagger
}

func (s *Config) Default() interface{} {
	return &Config{
		Enable: false,
		Auto:   true,
	}
}
