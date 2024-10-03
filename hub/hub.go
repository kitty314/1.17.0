package hub

import (
	"github.com/kitty314/1.17.0/config"
	"github.com/kitty314/1.17.0/hub/executor"
	"github.com/kitty314/1.17.0/hub/route"
	"github.com/kitty314/1.17.0/log"
)

type Option func(*config.Config)

func WithExternalUI(externalUI string) Option {
	return func(cfg *config.Config) {
		cfg.General.ExternalUI = externalUI
	}
}

func WithExternalController(externalController string) Option {
	return func(cfg *config.Config) {
		cfg.General.ExternalController = externalController
	}
}

func WithSecret(secret string) Option {
	return func(cfg *config.Config) {
		cfg.General.Secret = secret
	}
}

// Parse call at the beginning of clash.meta
func Parse(options ...Option) error {
	cfg, err := executor.Parse()
	if err != nil {
		return err
	}

	for _, option := range options {
		option(cfg)
	}

	if cfg.General.ExternalUI != "" {
		route.SetUIPath(cfg.General.ExternalUI)
	}

	if cfg.General.ExternalController != "" {
		go route.Start(cfg.General.ExternalController, cfg.General.ExternalControllerTLS,
			cfg.General.Secret, cfg.TLS.Certificate, cfg.TLS.PrivateKey, cfg.General.LogLevel == log.DEBUG)
	}

	executor.ApplyConfig(cfg, true)
	return nil
}
