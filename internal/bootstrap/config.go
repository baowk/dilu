package bootstrap

import (
	"dilu/common/config"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

func LoadConfig(configPath string) (*config.Config, error) {
	if strings.TrimSpace(configPath) == "" {
		return nil, fmt.Errorf("config path is required")
	}

	v := viper.New()
	v.SetConfigFile(configPath)
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("fatal error config file: %w", err)
	}

	var cfg config.Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config failed: %w", err)
	}

	if cfg.Server.RemoteEnable {
		if err := loadRemoteConfig(&cfg); err != nil {
			return nil, err
		}
		return &cfg, nil
	}

	config.SaveConfig(&cfg, nil)
	watchLocalConfig(v, &cfg)
	return &cfg, nil
}

func loadRemoteConfig(cfg *config.Config) error {
	rviper := viper.New()
	var err error
	if cfg.Remote.SecretKeyring == "" {
		err = rviper.AddRemoteProvider(cfg.Remote.Provider, cfg.Remote.Endpoint, cfg.Remote.Path)
	} else {
		err = rviper.AddSecureRemoteProvider(cfg.Remote.Provider, cfg.Remote.Endpoint, cfg.Remote.Path, cfg.Remote.SecretKeyring)
	}
	if err != nil {
		return fmt.Errorf("fatal error remote config provider: %w", err)
	}

	rviper.SetConfigType(cfg.Remote.GetConfigType())
	if err = rviper.ReadRemoteConfig(); err != nil {
		return fmt.Errorf("fatal error remote config read: %w", err)
	}

	var remoteCfg config.Config
	if err = rviper.Unmarshal(&remoteCfg); err != nil {
		return fmt.Errorf("unmarshal remote config failed: %w", err)
	}
	config.SaveConfig(cfg, &remoteCfg)

	go func() {
		for {
			time.Sleep(5 * time.Second)
			if watchErr := rviper.WatchRemoteConfig(); watchErr != nil {
				slog.Warn("watch remote config failed", "err", watchErr)
				continue
			}
			if unmarshalErr := rviper.Unmarshal(&remoteCfg); unmarshalErr != nil {
				slog.Warn("unmarshal remote config failed", "err", unmarshalErr)
				continue
			}
			config.SaveConfig(cfg, &remoteCfg)
		}
	}()
	return nil
}

func watchLocalConfig(v *viper.Viper, cfg *config.Config) {
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		if err := v.Unmarshal(cfg); err != nil {
			slog.Warn("unmarshal changed config failed", "err", err)
			return
		}
		config.SaveConfig(cfg, nil)
		slog.Info("config file changed", "file", e.Name)
	})
}
