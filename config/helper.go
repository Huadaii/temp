package config

import (
	"log"
	"os"
	"strconv"

	"github.com/gomodule/redigo/redis"
	"github.com/widuu/goini"
)

var Redis *Cacher
var dataPath = "./data"

//init 初始化redis
func init() {
	var err error
	var appconf *goini.Config
	//|| os.Getenv("yondor_dev") == "dev"
	if os.Getenv("OS") == "Windows_NT" {
		log.Println("内网配置")
		//本地环境
		appconf = goini.SetConfig(dataPath + "/conf/app-dev.ini")
		/*} else if os.Getenv("yondor_dev") == "dev" {
		//内网
		appconf = goini.SetConfig(dataPath + "/conf/app-dev.ini")*/
	} else {
		appconf = goini.SetConfig(dataPath + "/conf/app.ini")
	}
	host := appconf.GetValue("pedu-applets-redis", "host")
	port := appconf.GetValue("pedu-applets-redis", "port")
	password := appconf.GetValue("pedu-applets-redis", "password")
	IdleTimeout, _ := strconv.Atoi(appconf.GetValue("pedu-applets-redis", "IdleTimeout"))
	MaxIdle, _ := strconv.Atoi(appconf.GetValue("pedu-applets-redis", "MaxIdle"))
	/* 	host := Resource().String("pedu-schedule-redis.host")
	   	port := Resource().String("pedu-schedule-redis.port")
	   	password := Resource().String("pedu-schedule-redis.password")
	   	IdleTimeout, _ := Resource().Int("pedu-schedule-redis.IdleTimeout")
	   	MaxIdle, _ := Resource().Int("pedu-schedule-redis.MaxIdle") */
	log.Println("pedu-applets-redis:", host, port, IdleTimeout, MaxIdle)
	Redis, err = New(Options{
		Addr:        host + ":" + port,
		Password:    password,
		IdleTimeout: IdleTimeout,
		MaxIdle:     MaxIdle,
	})
	if err != nil {
		log.Println("pedu-applets-redis:", err)
		panic(err)
	}
}

// Int is a helper that converts a command reply to an integer
func Int(reply interface{}, err error) (int, error) {
	return redis.Int(reply, err)
}

// Int64 is a helper that converts a command reply to 64 bit integer
func Int64(reply interface{}, err error) (int64, error) {
	return redis.Int64(reply, err)
}

// String is a helper that converts a command reply to a string
func String(reply interface{}, err error) (string, error) {
	return redis.String(reply, err)
}

// String is a helper that converts a command reply to a string
func Strings(reply interface{}, err error) ([]string, error) {
	return redis.Strings(reply, err)
}

// Bool is a helper that converts a command reply to a boolean
func Bool(reply interface{}, err error) (bool, error) {
	return redis.Bool(reply, err)
}
