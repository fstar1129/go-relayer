package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/LATOKEN/relayer-smart-contract.git/src/app"
	"github.com/LATOKEN/relayer-smart-contract.git/src/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.NewViperConfig()
	srvURL := cfg.ReadServiceConfig()
	laCfg := cfg.ReadLachainConfig()
	chainCfgs := cfg.ReadWorkersConfig()
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
		panic(fmt.Sprintf("logrus err: %s", err.Error()))
	}
	logger.SetLevel(level)

	os.MkdirAll("./logs", os.ModePerm)
	logFile, err := os.OpenFile("./logs/relayer.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	logger.SetOutput(logFile)

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
	app := app.NewApp(logger, srvURL, db, laCfg, chainCfgs, resourceIDs)
	//run App
	app.Run(ctx)
}
