package configs

import (
	"fmt"
	"gonstructor/internal/sources"
	"time"

	"github.com/spf13/viper"
)

type AppConfig struct {
	LogLevel              string            `json:"log_level" bson:"log_level" yaml:"log_level" mapstructure:"log_level"`
	WebAddr               string            `json:"web_addr" bson:"web_addr" yaml:"web_addr" mapstructure:"web_addr"`
	ShutdownTimeout       time.Duration     `json:",omitempty"`
	ShutdownTimeoutString string            `json:"shutdown_timeout" bson:"shutdown_timeout" yaml:"shutdown_timeout" mapstructure:"shutdown_timeout"`
	RoseDB                *RoseDBConfig     `json:"rosedb" bson:"rosedb" yaml:"rosedb" mapstructure:"rosedb"`
	TG                    *sources.TGConfig `json:"tg" bson:"tg" yaml:"tg" mapstructure:"tg"`
}

func NewAppConfig(path string) (*AppConfig, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(path)

	err := v.ReadInConfig()

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	config := &AppConfig{}
	err = v.Unmarshal(config)

	config.ShutdownTimeout = v.GetDuration("shutdown_timeout")

	if err != nil {
		return nil, err
	}

	return config, nil
}
