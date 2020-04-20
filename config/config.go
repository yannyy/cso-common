package config

import (
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/env"
	"github.com/micro/go-plugins/config/source/configmap"
	"github.com/micro/v2/config/source/file"
)

// NewConfig returns config with env and k8s configmap setup
func NewConfig(opts ...config.Option, file_path string) config.Config {
	cfg := config.NewConfig()
	if file_path != "" {
		cfg.Load(
			env.NewSource(),
			configmap.NewSource(),
			file.NewSource(file_path),
		)
	} else {
		cfg.Load(
			env.NewSource(),
			configmap.NewSource(),
		)

	}
	return cfg
}
