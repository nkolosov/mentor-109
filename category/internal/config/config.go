package config

// Config конфигурация приложения
type Config struct {
	LogLevel string `long:"log-level" description:"Log level: panic, fatal, warn or warning, info, debug" env:"CL_LOG_LEVEL" required:"true"`
	LogJSON  bool   `long:"log-json" description:"Enable force log format JSON" env:"CL_LOG_JSON"`
	// HttpServiceListen хост и порт, который будет слушать HTTP сервер (в формате host:port)
	HttpServiceListen string `long:"http-service-listen" description:"Listening host:port for service http-server" env:"CL_HTTP_SERVICE_LISTEN" required:"true"`
	GrpcListen        string `long:"grpc-listen" description:"Listening host:port for grpc-server" env:"CL_GRPC_LISTEN" required:"true"`

	// EnablePprof включение отладки при помощи pprof
	EnablePprof bool `long:"enable-pprof" description:"Enable pprof server" env:"CL_ENABLE_PPROF"`

	DBHost     string `long:"db-host" description:"Database Server" env:"CL_DB_HOST" required:"true"`
	DBPort     int    `long:"db-port" description:"Database Port" env:"CL_DB_PORT" required:"true"`
	DBName     string `long:"db-name" description:"Database Name" env:"CL_DB_NAME" required:"true"`
	DBUser     string `long:"db-user" description:"Database User" env:"CL_DB_USER" required:"true"`
	DBPassword string `long:"db-pass" description:"Database Password" env:"CL_DB_PASSWORD" required:"true"`
}
