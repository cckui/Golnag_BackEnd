package _Log

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// ==========================
// @@	LogOutLevel=1	Log Write File & Stdout
// @@	LogOutLevel=2	Log Write Stdout
// @@	LogOutLevel=3	Log Write File
// @@	Format = 1	Normal
// @@ 	Format = 2	JSON
// @@	Path =	./logs/
// ==========================
type ModuleCfgStruct struct {
	OutLevel uint8  `json:"OutLevel"`
	Format      uint8  `json:"Format"`
	Path        string `json:"Path"`
}

var ModuleCfg *ModuleCfgStruct

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
func NewLogger(filePath string, level zapcore.Level, maxSize int, maxBackups int, maxAge int, compress bool, serviceName string) *zap.Logger {
	core := newCore(filePath, level, maxSize, maxBackups, maxAge, compress)
	return zap.New(core, zap.AddCaller(), zap.Development(), zap.Fields(zap.String("module", serviceName)))
}

/**
 * zapcore構造
 */
func newCore(filePath string, level zapcore.Level, maxSize int, maxBackups int, maxAge int, compress bool) zapcore.Core {
	//日誌文件路徑配置2
	hook := lumberjack.Logger{
		Filename:   filePath,   // 日誌文件路徑
		MaxSize:    maxSize,    // 每個日誌文件保存的最大尺寸 單位：M
		MaxBackups: maxBackups, // 日誌文件最多保存多少個備份
		MaxAge:     maxAge,     // 文件最多保存多少天
		Compress:   compress,   // 是否壓縮
	}
	// 設置日誌級別
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)
	//公用編碼器
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:  "time",
		LevelKey: "level",
		NameKey:  "logger",
		// CallerKey:     "path", // linenum
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder, // 小寫編碼器
		// EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 時間格式
		EncodeTime:     TimeEncoder,                    // 自訂 時間格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 路徑編碼器	FullCallerEncoder or ShortCallerEncoder
		EncodeName:     zapcore.FullNameEncoder,        // FullNameEncoder
	}
	switch ModuleCfg.Format {
	case 2: //JSON Format
		switch ModuleCfg.OutLevel {

		case 2: //Write	Stdout
			return zapcore.NewCore(
				zapcore.NewJSONEncoder(encoderConfig),                   // Json輸出(編碼器配置)
				zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), // 打印到控制台
				atomicLevel, // 日誌級別
			)
		case 3: //Write	File
			return zapcore.NewCore(
				zapcore.NewJSONEncoder(encoderConfig),               // Json輸出(編碼器配置)
				zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook)), // 打印到文件
				atomicLevel, // 日誌級別
			)
		default: //Write File && Stdout
			return zapcore.NewCore(
				zapcore.NewJSONEncoder(encoderConfig),                                           // Json輸出(編碼器配置)
				zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
				atomicLevel, // 日誌級別
			)
		}

	default: //Normal Format
		switch ModuleCfg.OutLevel {

		case 2: //Write	Stdout
			return zapcore.NewCore(
				zapcore.NewConsoleEncoder(encoderConfig),                // 一般輸出(編碼器配置)
				zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), // 打印到控制台
				atomicLevel, // 日誌級別
			)
		case 3: //Write	File
			return zapcore.NewCore(
				zapcore.NewConsoleEncoder(encoderConfig),            // 一般輸出(編碼器配置)
				zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook)), // 打印到文件
				atomicLevel, // 日誌級別
			)
		default: //Write File && Stdout
			return zapcore.NewCore(
				zapcore.NewConsoleEncoder(encoderConfig),                                        // 一般輸出(編碼器配置)
				zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
				atomicLevel, // 日誌級別
			)
		}
	}
}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	// enc.AppendString(t.Format(time.RFC3339)) //time.RFC3339
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}
