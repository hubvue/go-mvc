package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"server-test/src/config"
	"strings"
	"sync"
)

var (
	Mysql *sql.DB
	mu  sync.RWMutex
)

func NewMysql() *sql.DB {
	mu.RLock()
	if Mysql != nil {
		mu.RUnlock()
		return Mysql
	}
	mu.RUnlock()
	mu.Lock()
	path := strings.Join([]string{
		config.Conf.User,
		":",
		config.Conf.Password,
		"@tcp(",
		config.Conf.Mysql.Host,
		":",
		config.Conf.Mysql.Port,
		")/",
		config.Conf.Database,
		"?charset=utf8",
	}, "")
	fmt.Println(path)
	// 打开数据库，前者是驱动名，所以要导入：_ "github.com/go-sql-driver/mysql"
	Mysql, err := sql.Open("mysql", path)
	if err != nil {
		fmt.Println(err)
	}
	// 设置数据库最大连接数
	Mysql.SetMaxOpenConns(config.Conf.MaxOpen)
	// 设置数据库最大空闲连接数
	Mysql.SetMaxIdleConns(config.Conf.MaxIdle)
	if err := Mysql.Ping(); err != nil {
		fmt.Println("open database fail")
		return nil
	}
	fmt.Println("connect success")
	mu.Unlock()
	return Mysql
}

func CloseMysql() {
	if Mysql != nil {
		Mysql.Close()
	}
}
