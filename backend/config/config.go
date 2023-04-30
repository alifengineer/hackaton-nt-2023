package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const (
	// DebugMode indicates service mode is debug.
	DebugMode = DEBUG_ENVIRONMENT
	// TestMode indicates service mode is test.
	TestMode = TEST_ENVIRONMENT
	// ReleaseMode indicates service mode is release.
	ReleaseMode = PRODUCTION_ENVIRONMENT
)

type Config struct {
	ServiceName string
	Environment string // debug, test, release
	Version     string

	HTTPPort   string
	HTTPScheme string

	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string

	PostgresMaxConnections int32
	DefaultOffset          string
	DefaultLimit           string

	GoogleClientID         string
	GoogleClientSecret     string
	GoogleOAuthRedirectUrl string

	PlayMobileLogin      string
	PlayMobilePassword   string
	PlayMobileUrl        string
	PlayMobileOriginator string

	MinioEndpoint        string
	MinioAccessKeyID     string
	MinioSecretAccessKey string
}

// Load ...
func Load() Config {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.ServiceName = cast.ToString(getOrReturnDefaultValue("SERVICE_NAME", "ONLINE_BANKING"))
	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", DebugMode))
	config.Version = cast.ToString(getOrReturnDefaultValue("VERSION", "1.0"))

	config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":3000"))
	config.HTTPScheme = cast.ToString(getOrReturnDefaultValue("HTTP_SCHEME", "http"))

	config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "0.0.0.0"))
	config.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5454))
	config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "postgres"))
	config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "admin"))
	config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "xakaton_nt"))

	config.PostgresMaxConnections = cast.ToInt32(getOrReturnDefaultValue("POSTGRES_MAX_CONNECTIONS", 30))

	config.DefaultOffset = cast.ToString(getOrReturnDefaultValue("DEFAULT_OFFSET", "0"))
	config.DefaultLimit = cast.ToString(getOrReturnDefaultValue("DEFAULT_LIMIT", "15"))

	config.GoogleClientID = cast.ToString(getOrReturnDefaultValue("GOOGLE_OAUTH_CLIENT_ID", ""))
	config.GoogleClientID = cast.ToString(getOrReturnDefaultValue("GOOGLE_OAUTH_CLIENT_SECRET", ""))
	config.GoogleClientID = cast.ToString(getOrReturnDefaultValue("GOOGLE_OAUTH_REDIRECT_URL", ""))

	config.PlayMobileLogin = cast.ToString(getOrReturnDefaultValue("PLAY_MOBILE_LOGIN", "delever"))
	config.PlayMobilePassword = cast.ToString(getOrReturnDefaultValue("PLAY_MOBILE_PASSWORD", "vb6J9G6k2X"))
	config.PlayMobileUrl = cast.ToString(getOrReturnDefaultValue("PLAY_MOBILE_URL", "http://91.204.239.44/broker-api/send"))
	config.PlayMobileOriginator = cast.ToString(getOrReturnDefaultValue("PLAY_MOBILE_ORIGINATOR", "3700"))

	config.MinioAccessKeyID = cast.ToString(getOrReturnDefaultValue("MINIO_ACCESS_KEY", "ongei0upha4DiaThioja6aip8dolai1o"))
	config.MinioSecretAccessKey = cast.ToString(getOrReturnDefaultValue("MINIO_SECRET_KEY", "aew8aeheungohf7vaiphoh7Tusie2vei"))
	config.MinioEndpoint = cast.ToString(getOrReturnDefaultValue("MINIO_ENDPOINT", "test.cdn.u-code.io"))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
