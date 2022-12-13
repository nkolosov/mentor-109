package main

import (
	"database/sql"
	"fmt"
	"github.com/jessevdk/go-flags"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

const (
	host     = "0.0.0.0"
	port     = 5432
	user     = "username"
	password = "password"
	dbname   = "default_database"
)

func main() {
	var cfg Config
	parser := flags.NewParser(&cfg, flags.Default)
	_, err := parser.Parse()
	if err != nil {
		log.Fatal("Failed to parse config.", err)
	}

	logger, err := initLogger(cfg.LogLevel, cfg.LogJSON)
	if err != nil {
		log.Fatal("Failed to init logger.", err)
	}

	logger.Info("config", zap.Any("logger", cfg))

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	rows, err := db.Query(`SELECT city_name, country_id FROM city`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var name string
		var roll int

		err = rows.Scan(&name, &roll)
		if err != nil {
			panic(err)
		}

		fmt.Println(name, roll)
	}

	fmt.Println("Connected!")
}

// initLogger создает и настраивает новый экземпляр логгера
func initLogger(logLevel string, isLogJson bool) (*zap.Logger, error) {
	lvl := zap.InfoLevel
	err := lvl.UnmarshalText([]byte(logLevel))
	if err != nil {
		return nil, fmt.Errorf("can't unmarshal log-level: %w", err)
	}
	opts := zap.NewProductionConfig()
	opts.Level = zap.NewAtomicLevelAt(lvl)
	opts.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	if opts.InitialFields == nil {
		opts.InitialFields = map[string]interface{}{}
	}
	//opts.InitialFields["version"] = Version
	if !isLogJson {
		opts.Encoding = "console"
		opts.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	return opts.Build()
}
