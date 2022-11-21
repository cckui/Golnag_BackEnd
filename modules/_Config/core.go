package _Config

import (
	"github.com/spf13/viper"

	_Gin "project/modules/_Gin"
	api "project/modules/_Gin/controllers/api"
	_Log "project/modules/_Log"
	_MySQL "project/modules/_MySQL"
)

type AllConfigStruct struct {
	Mode     string                 `json:"Mode"`
	IP       string                 `json:"IP"`
	RemoteIP string                 `json:"RemoteIP"`
	CacheIP  string                 `json:"CacheIP"`
	Database _MySQL.ModuleCfgStruct `json:"Database"`
	MysqlA   MysqlConfig            `json:"DB_A"`
	MysqlS   MysqlConfig            `json:"DB_S"`
	Log      _Log.ModuleCfgStruct   `json:"Log"`
}

type MysqlConfig struct {
	IP       string `json:"ip"`
	Port     string `json:"port"`
	Account  string `json:"account"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
}

var AllConfig *AllConfigStruct

func ConfigInit() *AllConfigStruct {

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("json")
	viper.ReadInConfig()

	viper.SetDefault("Mode", "r")
	viper.SetDefault("IP", ":80")
	viper.SetDefault("RemoteIP", ":9090")
	viper.SetDefault("CacheIP", "127.0.0.1:6379")

	viper.SetDefault("Database.SQL_IP", "127.0.0.1:3306")
	viper.SetDefault("Database.SQL_Account", "root")
	viper.SetDefault("Database.SQL_Password", "123")

	viper.SetDefault("MysqlA.IP", "127.0.0.1")
	viper.SetDefault("MysqlA.Port", "3306")
	viper.SetDefault("MysqlA.Account", "root")
	viper.SetDefault("MysqlA.Password", "123")
	viper.SetDefault("MysqlA.DBName", "test")

	viper.SetDefault("MysqlS.IP", "127.0.0.1")
	viper.SetDefault("MysqlS.Port", "3306")
	viper.SetDefault("MysqlS.Account", "root")
	viper.SetDefault("MysqlS.Password", "123")
	viper.SetDefault("MysqlS.DBName", "test")

	viper.SetDefault("Log.OutLevel", 1)
	viper.SetDefault("Log.Format", 1)
	viper.SetDefault("Log.Path", "./logs/")

	err := viper.Unmarshal(&AllConfig)
	if err != nil {
		panic("Config Read Error")
		// log.Fatal(err)
	}

	_Gin.ModuleCfg.Mode = AllConfig.Mode
	_Gin.ModuleCfg.IP = AllConfig.IP

	// _Redis.Redis_IP = AllConfig.CacheIP

	// _MySQL.ModuleCfg = &AllConfig.Database

	_Log.ModuleCfg = &AllConfig.Log

	api.SysInfoInit(
		AllConfig.IP,
		AllConfig.Mode,
		AllConfig.RemoteIP,
		AllConfig.Database.SQL_IP)

	// fmt.Println(AllConfig)

	return AllConfig
}

// ===========================================================
