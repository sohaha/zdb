package zdb_test

import (
	"errors"
	"time"

	"github.com/sohaha/zdb"
	"github.com/sohaha/zdb/Driver/mssql"
	"github.com/sohaha/zdb/Driver/mysql"
	"github.com/sohaha/zdb/Driver/sqlite3"
	"github.com/sohaha/zlsgo/zutil"
)

var dbType = zutil.Getenv("dbtype", "sqlite3")

type TestTableUser struct {
	ID   int          `zdb:"id"`
	Name string       `zdb:"name"`
	Is   bool         `zdb:"is_ok"`
	Date zdb.JsonTime `zdb:"date"`
	Time time.Time
}

func (*TestTableUser) TableName() string {
	return "t_user"
}

var table *TestTableUser

func getDbConf(id string) (dbConf zdb.IfeConfig, err error) {
	switch dbType {
	case "sqlite3":
		dbConf = &sqlite3.Config{
			File:   "./test" + id + ".db",
			Memory: true,
		}
	case "mysql":
		dbConf = &mysql.Config{
			Host:     "127.0.0.1",
			Port:     33061,
			User:     "root",
			Password: "666666",
			DBName:   "test" + id,
		}
	case "mssql":
		dbConf = &mssql.Config{
			Dsn: "",
		}
	default:
		return nil, errors.New("未知数据库类型")
	}
	return
}

func initTable(db *zdb.Engine, dbType string) error {
	var err error
	switch dbType {
	case "sqlite3":
		_, _ = db.Exec(`DROP TABLE ` + table.TableName())
		_, err = db.Exec(`CREATE TABLE ` + table.TableName() + ` (id INTEGER NOT NULL,is_ok INTEGER,name TEXT,date DATE,time DATE,PRIMARY KEY (id));`)
	case "mysql":
		_, _ = db.Exec(`DROP TABLE ` + table.TableName())
		_, err = db.Exec(`CREATE TABLE ` + table.TableName() + ` (id int(0) NOT NULL AUTO_INCREMENT,is_ok int(1),name varchar(255) NULL,date datetime(0) NULL,PRIMARY KEY (id));`)
	case "mssql":
	}
	return err
}
