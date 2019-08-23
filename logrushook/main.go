package main

import (
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
)

func main() {
	// 24小时一个日志文件，最多存365天的日志文件。 再多了就删掉
	// ConfigLocalFilesystemLogger("log", "lg", time.Hour*24*365, time.Hour*24)
	ConfigLocalFilesystemLogger("code//logrus_test//log", "lg", time.Hour*24*365, time.Minute)
	ConfigLocalFilesystemLogger("code//logrus_test//log", "lg", time.Hour*24*365, time.Hour*24)

	for {
		time.Sleep(time.Second * 3)
		log.Info("hello ")
	}
}

// ConfigLocalFilesystemLogger config logrus log to local filesystem, with file rotation
func ConfigLocalFilesystemLogger(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) {
	baseLogPaht := path.Join(logPath, logFileName)
	writer, err := rotatelogs.New(
		baseLogPaht+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(baseLogPaht), // 生成软链，指向最新日志文件

		rotatelogs.WithMaxAge(maxAge), // 文件最大保存时间
		// rotatelogs.WithRotationCount(365),  // 最多存365个文件

		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	if err != nil {
		log.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer, // 为不同级别设置不同的输出目的
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, &log.TextFormatter{})
	log.AddHook(lfHook)
}
