package boot

import (
	"fmt"
	"uc/app/crontab"
	"uc/config"
	"uc/utils"
	"uc/utils/mylog"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func init() {
	//initConfig() // 初始化配置文件
	utils.InitNacos()
}

func Boot() {
	initZeroLogger() // 初始化日志功能

	utils.InitMysql()
	utils.InitRedis()

	go crontab.CronTab()
}

func initZeroLogger() {
	// 日志功能
	logConfig := mylog.Config{
		ConsoleLoggingEnabled: config.Cfg.Log.ConsoleLoggingEnabled,
		EncodeLogsAsJson:      config.Cfg.Log.EncodeLogsAsJson,
		FileLoggingEnabled:    config.Cfg.Log.FileLoggingEnabled,
		Directory:             config.Cfg.Log.Directory,
		Filename:              config.Cfg.Log.Filename,
		MaxSize:               config.Cfg.Log.MaxSize,
		MaxBackups:            config.Cfg.Log.MaxBackups,
		MaxAge:                config.Cfg.Log.MaxAge,
	}
	mylog.InitZeroLogger(logConfig)
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("config file is no found !！！"))
		}
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	conf := config.ConfigMap{}

	viper.OnConfigChange(func(e fsnotify.Event) {
		viper.Unmarshal(&conf)
		config.Cfg = &conf
		initZeroLogger()
	})
	viper.Unmarshal(&conf)
	viper.WatchConfig()
	config.Cfg = &conf
}
