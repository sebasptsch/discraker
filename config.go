package main

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

type MoonrakerConfigDefinition struct {
	ConnectionURL string `toml:"connection_url,commented" comment:"default url is the moonraker socket, can also be the url to your moonraker instance" `
}

type DiscordConfigDefinition struct {
	Token   string `toml:"token,commented" comment:"Fill this in with your discord bot token from https://discord.com/developers/applications"`
	GuildID string `toml:"guild_id,commented" comment:"Optionally add the ID of the server you want your bot to be in"`
}

type ConfigDefinition struct {
	Moonraker MoonrakerConfigDefinition `toml:"moonraker" comment:"Moonraker connection settings"`
	Discord   DiscordConfigDefinition   `toml:"discord" comment:"Discord connection settings"`
}

var Config = &ConfigDefinition{
	Moonraker: MoonrakerConfigDefinition{
		ConnectionURL: "xxx", // default values
	},
	Discord: DiscordConfigDefinition{
		Token:   "xxx", // default values
		GuildID: "xxx", // default values
	},
}

func expandPath(path string) string {
	if strings.HasPrefix(path, "~/") {
		home, err := os.UserHomeDir()
		if err == nil {
			return filepath.Join(home, path[2:])
		}
	}
	return path
}

func ReadConfig(path string) error {
	path = expandPath(path)
	_, err := os.Stat(path)
	if errors.Is(err, os.ErrPermission) {
		log.Fatalf("Path: %s is unwritable", path)
	} else if errors.Is(err, os.ErrNotExist) {
		out, err := toml.Marshal(Config)
		if err != nil {
			log.Fatalf("Config is not valid: %v", err)
		}
		err = os.WriteFile(path, out, 0644)

		if err != nil {
			log.Fatalf("Error writing default config %v", err)
		}
	}

	fileContents, err := os.ReadFile(path)

	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	err = toml.Unmarshal(fileContents, Config)
	if err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
	}

	return err
}
