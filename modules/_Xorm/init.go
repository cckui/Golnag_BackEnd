package sxorm

import (
	"fmt"

	_Log "project/modules/_Log"
	"strconv"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

type DbInfo struct {
	DbEngine   string
	DbIP       string
	DbPort     uint16
	DbAccount  string
	DbPassword string
	DbName     string
	Engine     *xorm.Engine
}

// dbEngine： mysql、postgres、mssql.
//
// dbIP：IP or sqlite3 path, port、dbAccount、dbPassword、dbName
func XormInit(dbEngine, dbIP string, args ...string) (*DbInfo, error) {

	var db *DbInfo

	if dbEngine == "sqlite3" {

		db = &DbInfo{
			DbEngine: dbEngine,
			DbIP:     dbIP,
		}
	} else if (dbEngine == "mysql" || dbEngine == "postgres" || dbEngine == "mssql") && len(args) == 4 {

		dbPortInt, _ := strconv.Atoi(args[0])
		dbPort := uint16(dbPortInt)
		db = &DbInfo{
			DbEngine:   dbEngine,
			DbIP:       dbIP,
			DbPort:     dbPort,
			DbAccount:  args[1],
			DbPassword: args[2],
			DbName:     args[3],
		}

	} else {
		return nil, fmt.Errorf("db args error")
	}

	var dbConnectInfo string

	switch db.DbEngine {

	case "mysql":
		//Format:"account:password(dbIP:dbPort)/dbName?charset=utf8"
		dbConnectInfo = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4",
			db.DbAccount, db.DbPassword, db.DbIP, db.DbPort, db.DbName)
	case "sqlite3":
		// ./test.db?cache=shared&mode=memory
		dbConnectInfo = fmt.Sprintf("%s?cache=shared&mode=memory", db.DbIP)
	case "postgres":
		dbConnectInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			db.DbIP, db.DbPort, db.DbAccount, db.DbPassword, db.DbName)
	case "mssql":
		dbConnectInfo = fmt.Sprintf("server=%s;port=%d;user id=%s;password=%s;database=%s;encrypt=disable",
			db.DbIP, db.DbPort, db.DbAccount, db.DbPassword, db.DbName)
	default:
		return nil, fmt.Errorf("db engine error")
	}

	en, err := xorm.NewEngine(db.DbEngine, dbConnectInfo)
	if err != nil {
		// fmt.Println("engine creation failed", err)
		_Log.MainLogger.Info("engine creation failed " + err.Error())
		return nil, err
	}

	//設定顯示Log
	en.ShowSQL(true)
	en.SetLogLevel(log.LOG_ERR) // core.LOG_ERR | core.LOG_DEBUG | core.LOG_INFO | core.LOG_OFF | core.LOG_UNKNOWN | core.LOG_WARNING

	//設定連線thread
	en.SetMaxOpenConns(5)
	en.SetMaxIdleConns(3)
	en.SetConnMaxLifetime(12 * time.Hour)

	// 設定快取
	// cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	// en.SetDefaultCacher(cacher)

	err = en.Ping()
	if err != nil {
		return nil, err
	}

	db.Engine = en

	_Log.MainLogger.Info(db.DbEngine + " 連線成功")

	// if err = en.Sync2(new(Account), new(User)); err != nil {
	// 	fmt.Println("Fail to sync database: %v\n", err)
	// }

	//ini table
	if err = en.Sync2(new(TbInfra)); err != nil {
		fmt.Println("Fail to sync database: %v\n", err)
		return nil, err
	}

	return db, nil
}
