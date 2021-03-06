package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/sohaha/zdb"
	"github.com/sohaha/zlsgo/zutil"
)

var _ zdb.IfeConfig = &Config{}

// Config database configuration
type Config struct {
	db         *sql.DB
	Dsn        string
	Host       string
	Port       int
	User       string
	Password   string
	DBName     string
	Parameters string
}

func (c *Config) DB() *sql.DB {
	db, _ := c.MustDB()
	return db
}

func (c *Config) MustDB() (*sql.DB, error) {
	var err error
	if c.db == nil {
		c.db, err = sql.Open(c.GetDriver(), c.GetDsn())
	}
	return c.db, err
}

func (c *Config) SetDB(db *sql.DB) {
	c.db = db
}

func (c *Config) GetDsn() string {
	if c.Dsn != "" {
		return c.Dsn
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		c.User, c.Password, c.Host, c.Port, c.DBName, zutil.IfVal(c.Parameters == "", "parseTime=true&charset=utf8&loc=Local", c.Parameters))

}

func (c *Config) GetDriver() string {
	return "mysql"
}
