package config

import (
	"io/ioutil"
	"path/filepath"

	"github.com/rs/zerolog"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Mad     MADConfig
	HA      HAConfig
	Devices []DeviceConfig
}

type MADConfig struct {
	Address  string
	Username string
	Password string
}

type HAConfig struct {
	Address string
}

type DeviceConfig struct {
	AdbAddress string
	Name       string
}

func ReadConfig() (*Config, error) {
	filename, _ := filepath.Abs("config.yaml")
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (c Config) MarshalZerologObject(e *zerolog.Event) {
	e.Object("mad", c.Mad).Object("ha", c.HA).Int("devices", len(c.Devices))
}

func (ha HAConfig) MarshalZerologObject(e *zerolog.Event) {
	e.Str("address", ha.Address)
}

func (m MADConfig) MarshalZerologObject(e *zerolog.Event) {
	password := ""
	if m.Password != "" {
		password = "********"
	}

	e.Str("address", m.Address).
		Str("username", m.Username).
		Str("password", password)
}
