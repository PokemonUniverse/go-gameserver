package main

import (
	"fmt"
	"runtime"

	"github.com/eaigner/hood"
	_ "github.com/ziutek/mymysql/godrv"

	"github.com/PokemonUniverse/nonamelib/configuration"
	"github.com/PokemonUniverse/nonamelib/configuration/providers"
	"github.com/PokemonUniverse/nonamelib/log"

	"gameserver/config"
	"gameserver/data/pokemon"
	"gameserver/server"
	"gameserver/world"
)

const (
	DEBUG     bool = true
	DEBUG_SQL bool = true
)

var mainDatabase *hood.Hood

func main() {
	// Always use the maximum available CPU/Cores
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Print header
	preInitialization()

	// TODO: Read configuration file
	readConfiguration()

	// Initialize database connection
	initializeDatabase()

	// Load all Pokemon data from database
	loadPokemonData()

	// Load tilepoints from database
	loadWorldmap()

	// Start listening for clients
	startGameServer()

	println("EXIT - End of code")
}

func preInitialization() {
	if DEBUG {
		log.Flags = log.L_CONSOLE
	} else {
		log.Flags = log.L_CONSOLE | log.L_FILE
		log.LogFilename = "puserver"
	}

	log.Println("*****************************************")
	log.Println("** Pokemon Universe Game Server	**")
	log.Println("** Version: 2.0			**")
	log.Println("** http://github.com/PokemonUniverse	**")
	log.Println("** GNU General Public License V2	**")
	log.Println("*****************************************")
}

func readConfiguration() {
	// Set configuration provider
	configuration.SetConfigurationProvider(providers.NewIniConfigProvider("serverconfig.ini"))
	// Add configuration items
	configuration.AddConfigurationItems(config.NewServerConfigItems())
	configuration.Initialize()

	log.Info("main", "readConfiguration", "Configuration initialized")
}

func initializeDatabase() {
	log.Verbose("main", "initializeDatabase", "Creating database connection")

	username, _ := configuration.GetString(config.CONFIG_DB_USER)
	password, _ := configuration.GetString(config.CONFIG_DB_PASSWORD)
	scheme, _ := configuration.GetString(config.CONFIG_DB_SCHEME)
	connectionString := fmt.Sprintf("%v/%v/%v", scheme, username, password)

	hd, err := hood.Open("mymysql", connectionString)
	if err != nil {
		log.Error("main", "initializeDatabase", "Unable to connect to database scheme '%s' with supplied credentials '%s'", scheme, username)
		panic("Unable to connect to database")
	}
	hd.Log = DEBUG_SQL

	mainDatabase = hd

	log.Info("main", "initializeDatabase", "Database connection initialized: %s", connectionString)
}

func loadPokemonData() {
	pokemon.Load(mainDatabase)
}

func loadWorldmap() {
	world.LoadWorldmap(mainDatabase)
}

func startGameServer() {
	port, err := configuration.GetInt(config.CONFIG_SERVER_PORT)
	if err != nil {
		port = 9001
		log.Error("main", "startGameServer", "Failed to read server port from configuration, using default 9001")
	}
	server.Listen(port)
}
