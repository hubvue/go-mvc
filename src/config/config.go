package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"sync"
)

type App struct {
	Host     string `ini:"host"`
	Port     string `ini:"port"`
	LogLevel string `ini:"log_level"`
}

type Mysql struct {
	User     string `ini:"user"`
	Host     string `ini:"host"`
	Port     string `ini:"port"`
	Password string `ini:"password"`
	Database string `ini:"database"`
	MaxIdle  int    `ini:"max_idle"`
	MaxOpen  int    `ini:"max_open"`
}

type Config struct {
	App
	Mysql
}

var (
	Conf *Config
	once sync.Once
)

func NewConfig(env string) {
	once.Do(func() {
		cfg, err := ini.ShadowLoad(fmt.Sprintf("src/config/%s.ini", env))
		if err != nil {
			log.Fatal(err)
		}
		config := new(Config)
		err = cfg.MapTo(config)
		Conf = config
	})
}
