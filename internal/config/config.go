package config

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
	Auth     AuthConfig
}

type ServerConfig struct {
	Host string
	Port int
}

type DatabaseConfig struct {
	URL string
}

type RedisConfig struct {
	URL string
}

type AuthConfig struct {
	AccessTokenTTL  int
	RefreshTokenTTL int
	RateLimitPerIP  int
	RateLimitPerAcct int
	JWKSPath        string
}
