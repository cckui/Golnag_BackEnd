package _MySQL

import (
	"database/sql"
	_Log "project/modules/_Log"

	_ "github.com/go-sql-driver/mysql"
)

type ModuleCfgStruct struct {
	SQL_IP       string `json:"SQL_IP"`
	SQL_Account  string `json:"SQL_Account"`
	SQL_Password string `json:"SQL_Password"`
}

var ModuleCfg *ModuleCfgStruct

var SqlDB *sql.DB

func DatabaseInit() {
	var err error
	SqlDB, err = sql.Open("mysql", ModuleCfg.SQL_Account+":"+ModuleCfg.SQL_Password+"@tcp("+ModuleCfg.SQL_IP+")/AGV?charset=utf8mb4,utf8&parseTime=false") //loc=local or &loc=Asia/Taipei
	if err != nil {
		_Log.DBLogger.Info(err.Error())
	}
	err = SqlDB.Ping()
	if err != nil {
		_Log.DBLogger.Info(err.Error())
		panic("Database Connect Error")
	}
}
