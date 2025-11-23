package main

import (
	"os"
	"path/filepath"

	toml "github.com/pelletier/go-toml/v2"
)

var (
	DefaultConf *Conf
	DefaultKey  *Key
)

func init() {
	DefaultConf = &Conf{}
	DefaultConf.load()
	DefaultKey = &Key{}
	DefaultKey.Load()
}

const (
	FCodeDir              = ".fcode"
	FCodeConfigFile       = "conf.toml"
	FCodeApiKeyFile       = "key.toml"
	FCodeCompletionPrompt = "!FCPR" + "EFIX!%s!FCSU" + "FFIX!%s!FCMI" + "DDLE!"
	IdeName               = "vim"
	PluginVersion         = "0.2.1"
	DefaultCursor         = "<CURSOR>"
	DefaultPort           = 8123
)

type Conf struct {
	UserName string `toml:"username"`
	Password string `toml:"password"`
	Cursor   string `toml:"cursor"`
	Port     int    `toml:"port"`
}

func (c *Conf) getPath() string {
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, FCodeDir, FCodeConfigFile)
}

func (c *Conf) load() {
	path := c.getPath()
	content, _ := os.ReadFile(path)
	if len(content) > 0 {
		_ = toml.Unmarshal(content, c)
	}
}

func (c *Conf) GetCursor() string {
	if c.Cursor != "" {
		return c.Cursor
	}
	return DefaultCursor
}

func (c *Conf) GetPort() int {
	if c.Port > 0 {
		return c.Port
	}
	return DefaultPort
}

type Key struct {
	APIKey string `toml:"key"`
}

func (k *Key) getPath() string {
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, FCodeDir, FCodeApiKeyFile)
}

func (k *Key) Load() {
	path := k.getPath()
	content, _ := os.ReadFile(path)
	if len(content) > 0 {
		_ = toml.Unmarshal(content, k)
	}
	if k.APIKey == "" {
		Login()
	}
}

func (k *Key) Save() {
	path := k.getPath()
	content, _ := toml.Marshal(k)
	os.WriteFile(path, content, os.ModePerm)
}
