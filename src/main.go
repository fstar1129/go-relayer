package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"gitlab.nekotal.tech/lachain/crosschain/relayer-smart-contract/src/app"
	"gitlab.nekotal.tech/lachain/crosschain/relayer-smart-contract/src/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.NewViperConfig()
	srvURL := cfg.ReadServiceConfig()
	laCfg := cfg.ReadLachainConfig()
	posCfg, bscCfg, ethCfg := cfg.ReadWorkersConfig()
	dbConfig := cfg.ReadDBConfig()
	dbURL := fmt.Sprintf(dbConfig.URL, dbConfig.DBHOST, dbConfig.DBPORT, dbConfig.DBUser, dbConfig.DBName, dbConfig.DBPassword, dbConfig.DBSSL)
	resourceIDs := cfg.ReadResourceIDs()
	// init logrus logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339Nano,
	})
	// set logger level
	level, err := logrus.ParseLevel(cfg.GetString("logger-level"))
	if err != nil {
		panic(err)
	}
	logger.SetLevel(level)

	// laworker := eth.NewErc20Worker(logger, laCfg)
	// laworker.GetBlockAndTxs(168078)

	// return

	// Set connection to onlife_business database
	db, err := gorm.Open(dbConfig.DBDriver, dbURL)
	if err != nil {
		logger.WithFields(logrus.Fields{"dbURL": dbURL}).Fatalf("Set connection to PostgreSQL: ", err.Error())
	}
	defer db.Close()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		sign := <-c
		logger.Infof("System signal: %+v\n", sign)
		cancel()
	}()
	app := app.NewApp(logger, srvURL, db, laCfg, posCfg, bscCfg, ethCfg, resourceIDs)
	//run App
	app.Run(ctx)
}
