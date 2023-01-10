package config

// Config конфигурация приложения
type Config struct {
	LogLevel          string `long:"log-level" description:"Log level: panic, fatal, warn or warning, info, debug" env:"CL_LOG_LEVEL" required:"true"`
	LogJSON           bool   `long:"log-json" description:"Enable force log format JSON" env:"CL_LOG_JSON"`
	HttpPrivateListen string `long:"http-private-listen" description:"Listening host:port for private http-server" env:"CL_HTTP_PRIVATE_LISTEN" required:"true"`
	GrpcListen        string `long:"grpc-listen" description:"Listening host:port for grpc-server" env:"CL_GRPC_LISTEN" required:"true"`

	DBHost     string `long:"db-host" description:"Database Server" env:"CL_DB_HOST" required:"true"`
	DBPort     int    `long:"db-port" description:"Database Port" env:"CL_DB_PORT" required:"true"`
	DBName     string `long:"db-name" description:"Database Name" env:"CL_DB_NAME" required:"true"`
	DBUser     string `long:"db-user" description:"Database User" env:"CL_DB_USER" required:"true"`
	DBPassword string `long:"db-pass" description:"Database Password" env:"CL_DB_PASSWORD" required:"true"`
}
