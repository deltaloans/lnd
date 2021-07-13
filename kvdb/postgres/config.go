package postgres

import "time"

// Config holds postgres configuration data.
type Config struct {
	Dsn     string        `long:"dsn" description:"Database connection string."`
	Timeout time.Duration `long:"timeout" description:"Database connection timeout. Set to zero to disable."`
}
