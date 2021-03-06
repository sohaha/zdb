package sqlite3

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/sohaha/zdb"
	"github.com/sohaha/zlsgo/zfile"
	"github.com/sohaha/zlsgo/zutil"
)

var _ zdb.IfeConfig = &Config{}

// Config database configuration
type Config struct {
	File   string
	Dsn    string
	db     *sql.DB
	Memory bool
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
	return "file:" + zfile.RealPath(c.File) + zutil.IfVal(c.Memory, "?cache=shared&mode=memory", "").(string)
}

func (c *Config) GetDriver() string {
	return "sqlite3"
}
