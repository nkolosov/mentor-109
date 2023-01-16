package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpczap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/jessevdk/go-flags"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	grpc_server "github.com/nkolosov/mentor-109/internal/api/grpc"
	"github.com/nkolosov/mentor-109/internal/config"
	"github.com/nkolosov/mentor-109/internal/repository/potgresql"
	categoryv1 "github.com/nkolosov/mentor-109/pkg/api/grpc/gen/auction/category/category/v1"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := initConfig()
	if err != nil {
		log.Fatal("Failed to init config.", err)
	}

	logger, err := initLogger(cfg.LogLevel, cfg.LogJSON)
	if err != nil {
		log.Fatal("Failed to init logger.", err)
	}

	logger.Info("config", zap.Any("logger", cfg))

	defer func() {
		if msg := recover(); msg != nil {
			err := fmt.Errorf("%s", msg)
			logger.Error("recovered from panic, but application will be terminated", zap.Error(err))
		}
	}()

	db, err := initDb(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	if err != nil {
		logger.Fatal("failed to initialize db", zap.Error(err))
	}

	defer func() {
		err := db.Close()
		if err != nil {
			logger.Error("failed to close db connection", zap.Error(err))
		}
	}()

	err = migrateUp(db)
	if err != nil {
		logger.Fatal("failed to migrations", zap.Error(err))
	}

	fmt.Println("Connected!")

	err = startGRPCServer(context.Background(), cfg.GrpcListen, db, logger)
	if err != nil {
		logger.Fatal("failed to start grpc server", zap.Error(err))
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}

func startGRPCServer(
	ctx context.Context,
	listen string,
	db *sqlx.DB,
	logger *zap.Logger,
) error {
	logger.Info("gRPC started", zap.String("listen", listen))
	lis, err := net.Listen("tcp", listen)
	if err != nil {
		return fmt.Errorf("failed to listen GRPC server: %v", err)
	}

	recoverFromPanicHandler := func(p interface{}) error {
		err := fmt.Errorf("recovered from panic: %s", p)
		logger.Error("recovered from panic", zap.Error(err))

		return err
	}
	opts := []grpcrecovery.Option{
		grpcrecovery.WithRecoveryHandler(recoverFromPanicHandler),
	}

	grpc_prometheus.EnableHandlingTimeHistogram()
	s := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpcrecovery.UnaryServerInterceptor(opts...),
			grpc_prometheus.UnaryServerInterceptor,
		),
		grpc_middleware.WithStreamServerChain(
			grpcrecovery.StreamServerInterceptor(opts...),
			grpczap.StreamServerInterceptor(logger),
			grpc_prometheus.StreamServerInterceptor,
		))

	categoryv1.RegisterCategoryAPIServer(
		s,
		grpc_server.NewCategoryServer(
			potgresql.NewCategoryRepository(db, logger),
			grpc_server.NewToProtobufMapper(),
		),
	)

	reflection.Register(s)
	return s.Serve(lis)
}

func initConfig() (config.Config, error) {
	var cfg config.Config
	parser := flags.NewParser(&cfg, flags.Default)
	_, err := parser.Parse()
	if err != nil {
		return cfg, fmt.Errorf("failed to parse config: %w", err)
	}

	return cfg, nil
}

func initDb(host string, port int, user string, password string, dbname string) (*sqlx.DB, error) {
	psqlconn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sqlx.Open("postgres", psqlconn)
	if err != nil {
		return nil, fmt.Errorf("failed to open db connection: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("ping to db failed: %w", err)
	}

	return db, nil
}

func migrateUp(db *sqlx.DB) error {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to get driver: %w", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://category/internal/migrations",
		"postgres", driver)
	if err != nil {
		return fmt.Errorf("failed to get migrations: %w", err)
	}
	err = m.Up()
	if err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("failed to apply migrations: %w", err)
		}
	}

	return nil
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
