package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/thrasher-corp/gocryptotrader/common"
	"github.com/thrasher-corp/gocryptotrader/config"
	"github.com/thrasher-corp/gocryptotrader/core"
	"github.com/thrasher-corp/gocryptotrader/database"
	dbPSQL "github.com/thrasher-corp/gocryptotrader/database/drivers/postgres"
	dbsqlite3 "github.com/thrasher-corp/gocryptotrader/database/drivers/sqlite3"
	"github.com/thrasher-corp/gocryptotrader/database/repository"
)

var (
	dbConn         *database.Instance
	configFile     string
	defaultDataDir string
	migrationDir   string
	command        string
	args           string
)

func openDBConnection(driver string) (err error) {
	if driver == database.DBPostgreSQL {
		dbConn, err = dbPSQL.Connect()
		if err != nil {
			return fmt.Errorf("database failed to connect: %v, some features that utilise a database will be unavailable", err)
		}
		return nil
	} else if driver == database.DBSQLite || driver == database.DBSQLite3 {
		dbConn, err = dbsqlite3.Connect()
		if err != nil {
			return fmt.Errorf("database failed to connect: %v, some features that utilise a database will be unavailable", err)
		}
		return nil
	}
	return errors.New("no connection established")
}

func main() {
	fmt.Println("GoCryptoTrader database seeding tool")
	fmt.Println(core.Copyright)
	fmt.Println()

	flag.StringVar(&command, "command", "", "command to run seed|status|up|up-by-one|up-to|down|create")
	flag.StringVar(&args, "args", "", "arguments to pass to goose")
	flag.StringVar(&configFile, "config", config.DefaultFilePath(), "config file to load")
	flag.StringVar(&defaultDataDir, "datadir", common.GetDefaultDataDir(runtime.GOOS), "default data directory for GoCryptoTrader files")
	flag.StringVar(&migrationDir, "migrationdir", database.MigrationDir, "override migration folder")

	flag.Parse()

	var conf config.Config
	err := conf.LoadConfig(configFile, true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if !conf.Database.Enabled {
		fmt.Println("Database support is disabled")
		os.Exit(1)
	}

	err = openDBConnection(conf.Database.Driver)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	drv := repository.GetSQLDialect()

	if drv == database.DBSQLite || drv == database.DBSQLite3 {
		fmt.Printf("Database file: %s\n", conf.Database.Database)
	} else {
		fmt.Printf("Connected to: %s\n", conf.Database.Host)
	}

	err = dbConn.SQL.Close()
	if err != nil {
		fmt.Println(err)
	}
}
