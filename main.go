package main

import (
	"embed"
	"flag"
	"fmt"
	_Config "project/modules/_Config"
	_Gin "project/modules/_Gin"
	_Log "project/modules/_Log"
	sxorm "project/modules/_Xorm"

	"github.com/google/uuid"

	// _MySQL "project/modules/_MySQL"
	// _Redis "project/modules/_Redis"

	_ "github.com/sijms/go-ora/v2"
)

//go:embed views/* public/*
var f embed.FS
var (
	username   = flag.String("uname", "scott", "oracle username")
	password   = flag.String("password", "tiger", "oracle password")
	oraclehost = flag.String("oraclehost", "dbhost", "oracle database host")
	oracleport = flag.Int("oracleport", 1521, "oracle database port")
	dbname     = flag.String("dbname", "orclpdb1", "oracle database name")
)

func main() {
	flag.Parse()

	//===== Load Config ./config.json
	config := _Config.ConfigInit()

	//===== Log
	_Log.Loginit()

	//===== Redis
	//_Redis.RedisInit()

	// +-----------------------------------
	// |	XORM
	// +-----------------------------------
	mydb0, err := sxorm.XormInit("mysql", config.MysqlA.IP, config.MysqlA.Port, config.MysqlA.Account, config.MysqlA.Password, config.MysqlA.DBName) //mysql
	// mydb1, err := sxorm.XormInit("mysql", config.MysqlS.IP, config.MysqlS.Port, config.MysqlS.Account, config.MysqlS.Password, config.MysqlS.DBName) //mysql

	if err != nil {
		panic(err)
	}

	go test(mydb0)

	//===== 設定Gin運行模式
	_Gin.GinInit(f)
}
func test(db *sxorm.DbInfo) {

	// +-------------------------------------
	// | 	單筆新增測試
	// +-------------------------------------
	var insertCount int64 = 0
	for i := 0; i < 100; i++ {

		uuid := uuid.New()
		key := uuid.String()
		tempName := fmt.Sprintf("test%d", i)
		// random, _ := rand.Int(rand.Reader, big.NewInt(2)) // 亂數產生0~100
		// insertInt, err := db.Engine.Insert(&User{Id: 111, Name: tempName, Sex: (i % 2) + 1})
		insertInt, err := db.Engine.Insert(&sxorm.TbInfra{Data1: key, Data2: tempName, Data3: i, Data4: true})

		insertCount += insertInt

		if err != nil {
			panic(err)
		}
	}
	fmt.Println(fmt.Sprintf("單筆新增100次：%d", insertCount))

}
