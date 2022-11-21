package _Log

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var MainLogger *zap.Logger
var WebLogger *zap.Logger
var DBLogger *zap.Logger
var WSLogger *zap.Logger

func Loginit() {

	if ModuleCfg.Path == "" {
		ModuleCfg.Path = "./logs/"
	}

	timeflag := time.Now().Format("2006-01-02_150405")
	MainLogger = NewLogger(ModuleCfg.Path+timeflag+"_main.log", zapcore.InfoLevel, 100, 10, 7, true, "Main")
	WebLogger = NewLogger(ModuleCfg.Path+timeflag+"_web.log", zapcore.InfoLevel, 100, 10, 7, true, "Web")
	DBLogger = NewLogger(ModuleCfg.Path+timeflag+"_db.log", zapcore.InfoLevel, 100, 10, 7, true, "DB")
	WSLogger = NewLogger(ModuleCfg.Path+timeflag+"_ws.log", zapcore.InfoLevel, 100, 10, 7, true, "ws")
}

/**
 * 獲取日誌
 * filePath 日誌文件路徑
 * level 日誌級別
 * maxSize 每個日誌文件保存的最大尺寸 單位：M
 * maxBackups 日誌文件最多保存多少個備份
 * maxAge 文件最多保存多少天
 * compress 是否壓縮
 * serviceName 服務名
 */

//  log.MainLogger.Info(fmt.Sprintf("| %3d | %13v | %15s | %s  %s |", statusCode, latency, clientIP, method, path),
// 				zap.Int("status", statusCode),
// 				zap.String("method", method),
// 				zap.String("path", path),
// 				zap.Duration("latency", latency),
// 				zap.String("ip", clientIP),
// 				zap.String("user-agent", c.Request.UserAgent()),
// 				zap.String("query", query),
// 			)
// 				zap.String("query", query),
// 			)
// 				zap.String("query", query),
// 			)
