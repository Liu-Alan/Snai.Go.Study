package main

/*
import (
	log "github.com/Sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}
*/
/*
import (
	"os"

	log "github.com/Sirupsen/logrus"
)

func init() {
	// 日志格式化为JSON而不是默认的ASCII
	log.SetFormatter(&log.JSONFormatter{})

	// 输出stdout而不是默认的stderr，也可以是一个文件
	log.SetOutput(os.Stdout)

	// 设置记录日志级别
	log.SetLevel(log.InfoLevel)
}

func main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")

	// 通过日志语句重用字段
	// logrus.Entry返回自WithFields()
	contextLogger := log.WithFields(log.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})

	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")
}
*/
/*
import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	// 日志格式化为JSON而不是默认的ASCII
	log.SetFormatter(&logrus.JSONFormatter{})

	// 输出stdout而不是默认的stderr，也可以是一个文件
	log.SetOutput(os.Stdout)

	// 设置记录日志级别
	log.SetLevel(logrus.InfoLevel)

	//设置输出到文件
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file")
	}
}

func main() {
	log.WithFields(logrus.Fields{
		"filename": "123.txt",
	}).Info("打开文件失败")
}
*/

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"time"
)

func newLfsHook(logName string, logLevel string) log.Hook {
	writer, err := rotatelogs.New(
		logName+".%Y%m%d%H",
		// WithLinkName为最新的日志建立软连接，以方便随着找到当前日志文件
		rotatelogs.WithLinkName(logName),

		// WithRotationTime设置日志分割的时间，这里设置为一天分割一次
		rotatelogs.WithRotationTime(time.Hour*24),

		// WithMaxAge和WithRotationCount二者只能设置一个，
		// WithMaxAge设置文件清理前的最长保存时间，
		// WithRotationCount设置文件清理前最多保存的个数。
		//rotatelogs.WithRotationCount(maxRemainCnt),
		rotatelogs.WithMaxAge(time.Hour*24*30),
	)

	if err != nil {
		log.Errorf("config local file system for logger error: %v", err)
	}

	// 日志格式化为JSON而不是默认的ASCII
	log.SetFormatter(&logrus.JSONFormatter{})

	// 设置记录日志级别
	level, err := log.ParseLevel(logLevel)
	if err != nil {
		log.SetLevel(log.WarnLevel)
	} else {
		log.SetLevel(level)
	}

	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer,
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, &log.JSONFormatter{})

	return lfsHook
}

func main() {
	logName := "log.txt"
	logLevel := "info"
	log.AddHook(newLfsHook(logName, logLevel))
	log.Info("测试日志写入")
	log.WithFields(log.Fields{
		"parms": "参数不对",
	}).Error("测试日志警告写入")
}
