package configs

type RoseDBConfig struct {
	Path string `json:"path" bson:"path" yaml:"path" mapstructure:"path"`
}
