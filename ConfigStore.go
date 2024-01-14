package wailsconfigstore

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/adrg/xdg"
)

type Config string
type ConfigStore struct {
	configPath string
}

func NewConfigStore(folderName string) (*ConfigStore, error) {
	configFilePath, err := xdg.ConfigFile(folderName)
	if err != nil {
		return nil, fmt.Errorf("could not resolve path for config file: %w", err)
	}

	return &ConfigStore{
		configPath: configFilePath,
	}, nil
}

func (s *ConfigStore) Get(fileName string, defaultValue string) (Config, error) {
	_, err := os.Stat(s.configPath + string(os.PathSeparator) + fileName)
	if os.IsNotExist(err) {
		return Config(defaultValue), nil
	}
	buf, err := fs.ReadFile(os.DirFS(s.configPath), fileName)
	if err != nil {
		return Config(defaultValue), fmt.Errorf("could not read the configuration file: %w", err)
	}
	return Config(buf), nil
}

func (s *ConfigStore) Set(fileName string, value Config) error {
	err := os.MkdirAll(s.configPath, 0755)
	if err != nil {
		return fmt.Errorf("could not create the configuration directory: %w", err)
	}
	err = os.WriteFile(s.configPath+string(os.PathSeparator)+fileName, []byte(value), 0644)
	if err != nil {
		return fmt.Errorf("could not write the configuration file: %w", err)
	}
	return nil
}
