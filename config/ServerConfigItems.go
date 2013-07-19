package config

import (
	"github.com/PokemonUniverse/nonamelib/configuration"
)

const (
	CONFIG_SERVER_PORT string = "server_port"
	CONFIG_DB_USER     string = "db_user"
	CONFIG_DB_PASSWORD string = "db_password"
	CONFIG_DB_SCHEME   string = "db_scheme"
)

type ServerConfigItems struct {
	configItems map[string]configuration.IConfigurationItem
}

func NewServerConfigItems() *ServerConfigItems {
	serverConfig := &ServerConfigItems{}
	serverConfig.configItems = map[string]configuration.IConfigurationItem{
		CONFIG_SERVER_PORT: configuration.NewConfigurationItem("server", "port", "Port on which the server listens to new connections", 9001),
		CONFIG_DB_USER:     configuration.NewConfigurationItem("database", "user", "Username with which to connect to the database", "root"),
		CONFIG_DB_PASSWORD: configuration.NewConfigurationItem("database", "password", "Password for database user", ""),
		CONFIG_DB_SCHEME:   configuration.NewConfigurationItem("database", "scheme", "Database scheme to load", "puserver2")}

	return serverConfig
}

func (s *ServerConfigItems) GetConfigurationItems() map[string]configuration.IConfigurationItem {
	return s.configItems
}
