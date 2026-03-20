package bootstrap

import (
	"context"
	"dilu/internal/common/config"
	"fmt"
	"strings"
	"time"

	"github.com/baowk/dilu-core/core/logger"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

// configCancel 用于停止远程配置监听 goroutine
var configCancel context.CancelFunc

func LoadConfig(configPath string) (*config.Config, error) {
	if strings.TrimSpace(configPath) == "" {
		return nil, fmt.Errorf("config path is required")
	}

	v := viper.New()
	v.SetConfigFile(configPath)
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("fatal error config file: %w", err)
	}

	// 支持通过环境变量覆盖敏感配置，优先级高于配置文件
	// 例：export JWT_SIGN_KEY=your-secret 即可覆盖 jwt.sign-key
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.AutomaticEnv()
	_ = v.BindEnv("jwt.sign-key", "JWT_SIGN_KEY")
	_ = v.BindEnv("dbcfg.dns", "DB_DNS")
	_ = v.BindEnv("cache.password", "REDIS_PASSWORD")

	var cfg config.Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config failed: %w", err)
	}

	if cfg.Server.RemoteEnable {
		ctx, cancel := context.WithCancel(context.Background())
		configCancel = cancel
		if err := loadRemoteConfig(ctx, &cfg); err != nil {
			cancel()
			return nil, err
		}
		return &cfg, nil
	}

	config.SaveConfig(&cfg, nil)
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	watchLocalConfig(v, &cfg)
	return &cfg, nil
}

// StopConfigWatch 停止远程配置监听，用于优雅退出
func StopConfigWatch() {
	if configCancel != nil {
		configCancel()
	}
}

func loadRemoteConfig(ctx context.Context, cfg *config.Config) error {
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
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				logger.Info("remote config watcher stopped")
				return
			case <-ticker.C:
				if watchErr := rviper.WatchRemoteConfig(); watchErr != nil {
					logger.Warn("watch remote config failed", "err", watchErr)
					continue
				}
				if unmarshalErr := rviper.Unmarshal(&remoteCfg); unmarshalErr != nil {
					logger.Warn("unmarshal remote config failed", "err", unmarshalErr)
					continue
				}
				config.SaveConfig(cfg, &remoteCfg)
			}
		}
	}()
	return nil
}

func watchLocalConfig(v *viper.Viper, cfg *config.Config) {
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		if err := v.Unmarshal(cfg); err != nil {
			logger.Warn("unmarshal changed config failed", "err", err)
			return
		}
		config.SaveConfig(cfg, nil)
		logger.Info("config file changed", "file", e.Name)
	})
}
