package api

import (
	"strconv"
	"time"
)

var AllCfg interface{}
var sysData infoJSON

func SysInfoInit(webPort, webMode, wsIP, dbIP string) {

	fullTime, timestamp := GetSysTime()

	sysData = infoJSON{
		Version:  app_ver,
		SysTime:  fullTime + " | " + timestamp,
		Web_Port: webPort,
		Web_Mode: webMode,
		WS_IP:    wsIP,
		DB_IP:    dbIP,
	}
}

func SysVer() string {

	return app_ver
}

func GetSysTime() (timeFullStr, timeUnixStr string) {

	timeUnix := time.Now().Unix()
	timeFullStr = time.Now().Format("2006-01-02 15:04:05")
	timeUnixStr = strconv.FormatInt(timeUnix, 10)

	return
}

func ReturnSysInfo() *infoJSON {

	fullTime, timestamp := GetSysTime()

	sysData.SysTime = fullTime + " | " + timestamp

	return &sysData
}
