package db

import (
	"fmt"
	"time"
)

// Config is used to configure a postgres sql server connection
type Settings struct {
	DBName             string
	SSLMode            string
	User               string
	Password           string
	Host               string
	Port               string
	MaxConnections     int
	MaxIdleConnections int
	MaxConnLifetime    time.Duration
}

// ConnStr returns a postgres connection string
func (c *Settings) String() string {
	return fmt.Sprintf("dbname=%s sslmode=%s user=%s password=%s host=%s port=%s",
		c.DBName,
		c.SSLMode,
		c.User,
		c.Password,
		c.Host,
		c.Port)
}
