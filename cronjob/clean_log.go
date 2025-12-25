package cronjob

import (
	"gin-mini-agent/pkg/global"
	"log/slog"
	"runtime/debug"
	"time"
)

// 清理超过一周的日志表数据
type CleanLog struct {
}

func (u CleanLog) Run() {
	startTime := time.Now()
	global.Log.Debug("cronjob定时任务:CleanLog开始执行", slog.String("startTime", startTime.Format("2006-01-02 15:04:05")))
	defer func() {
		if panicErr := recover(); panicErr != nil {
			global.Log.Error("cronjob定时任务:CleanLog执行失败", 
				slog.Any("error", panicErr),
				slog.String("stack", string(debug.Stack())),
			)
		}
	}()

}
