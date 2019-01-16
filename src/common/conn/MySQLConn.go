package conn

import (
	"../config"
	"../utils/LoggerUtils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var mysqlPool chan *sqlx.DB

func MySQLConn() *sqlx.DB {
	host, port, username, password, db := config.GetMysqlConfig()
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", username, password, host, port, db)
	return initMySQL(connStr)
}

func putMySQL(conn *sqlx.DB) {
	if mysqlPool == nil {
		mysqlPool = make(chan *sqlx.DB, maxPoolSize)
	}
	if len(mysqlPool) >= maxPoolSize {
		_ = conn.Close()
		return
	}
	mysqlPool <- conn
}

func initMySQL(connStr string) *sqlx.DB {
	if len(mysqlPool) == 0 {
		// 如果长度为0，就定义一个*sqlx.DB类型长度为maxPoolSize的channel
		mysqlPool = make(chan *sqlx.DB, maxPoolSize)
		go func() {
			for i := 0; i < maxPoolSize/2; i++ {
				conn, err := sqlx.Open("mysql", connStr)
				if conn != nil {
					putMySQL(conn)
				} else {
					LoggerUtils.Error(err)
				}
			}
		}()
	}
	return <- mysqlPool
}
