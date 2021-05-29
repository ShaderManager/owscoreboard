package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Port          int    `yaml:"port"`
	TwitchChannel string `yaml:"twitch-channel"`

	Keymaps struct {
		IncTankWins   string `yaml:"increment-tank-wins"`
		IncTankTies   string `yaml:"increment-tank-ties"`
		IncTankLosses string `yaml:"increment-tank-losses"`
		IncDpsWins    string `yaml:"increment-dps-wins"`
		IncDpsTies    string `yaml:"increment-dps-ties"`
		IncDpsLosses  string `yaml:"increment-dps-losses"`
		IncSupWins    string `yaml:"increment-sup-wins"`
		IncSupTies    string `yaml:"increment-sup-ties"`
		IncSupLosses  string `yaml:"increment-sup-losses"`
		ResetStats    string `yaml:"reset-stats"`
	} `yaml:"keymaps"`
}

func NewConfig() Config {
	res := Config{}
	res.Port = 8080

	// Initialize default values
	res.Keymaps.IncTankWins = "numpad 7"
	res.Keymaps.IncTankTies = "numpad 8"
	res.Keymaps.IncTankLosses = "numpad 9"
	res.Keymaps.IncDpsWins = "numpad 4"
	res.Keymaps.IncDpsTies = "numpad 5"
	res.Keymaps.IncDpsLosses = "numpad 6"
	res.Keymaps.IncSupWins = "numpad 1"
	res.Keymaps.IncSupTies = "numpad 2"
	res.Keymaps.IncSupLosses = "numpad 3"
	res.Keymaps.ResetStats = "numpad 0"

	return res
}

func loadConfig(path string) (Config, error) {
	f, err := os.Open(path)
	cfg := NewConfig()

	if err != nil {
		return cfg, err
	}

	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)

	if err != nil {
		log.Printf("Failed to parse %s: %+v", path, err)
	}

	return cfg, err
}

var cfg Config

func init() {
	cfg = NewConfig()
}

func main() {
	cfg, _ = loadConfig("config.yaml")

	go startTwitchPolling()
	go setupKeyboardHook()

	startWebServer()
}
