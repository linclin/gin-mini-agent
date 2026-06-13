package initialize

import (
	"fmt"
	"gin-mini-agent/pkg/global"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"runtime/debug"

	"github.com/natefinch/lumberjack"
)

// Logger 初始化日志
func Logger() {
	// 创建日志目录
	if err := os.MkdirAll(global.Conf.Logs.Path, 0755); err != nil {
		panic(fmt.Sprintf("创建日志目录失败: %v", err))
	}

	// 使用 lumberjack 进行日志轮转
	logFile := &lumberjack.Logger{
		Filename:   filepath.Join(global.Conf.Logs.Path, global.Conf.System.AppName+".log"),
		MaxSize:    global.Conf.Logs.MaxSize,
		MaxBackups: global.Conf.Logs.MaxBackups,
		MaxAge:     global.Conf.Logs.MaxAge,
		Compress:   global.Conf.Logs.Compress,
	}

	// 同时输出到控制台和文件
	handler := slog.NewJSONHandler(io.MultiWriter(os.Stdout, logFile), &slog.HandlerOptions{
		AddSource: true,
		Level:     global.Conf.Logs.Level,
	})
	logger := slog.New(handler)
	slog.SetDefault(logger)
	global.Log = logger
	global.Log.Info("初始化日志完成")

	// panic 日志
	panicFile, err := os.Create(filepath.Join(global.Conf.Logs.Path, "panic.log"))
	if err != nil {
		global.Log.Error("初始化panic日志失败", slog.String("error", err.Error()))
		panic(err)
	}
	debug.SetCrashOutput(panicFile, debug.CrashOptions{})
	global.Log.Info("初始化panic日志完成")
}
