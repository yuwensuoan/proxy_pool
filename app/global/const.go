package global

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"proxy_pool/config"
	"time"
)

var Logger *logrus.Logger

const CodeSuccess = 0         // 处理成功
const ErrSystem = 50000       // 系统错误
const ErrInvalidParam = 10000 // 参数错误


// 获取code对应的信息
func GetMsgByCode(code int) string {
	m := map[int]string{
		CodeSuccess:     "Success",
		ErrSystem:       "System Error",
		ErrInvalidParam: "Parameter Error",
	}

	msg, ok := m[code]
	if !ok {
		msg = "System Error!"
	}
	return msg
}



func init()  {
	logFilePath := config.CONFIG.Logger.Filepath
	logFileName := config.CONFIG.Logger.Filename

	// 日志文件
	fileName := path.Join(logFilePath, logFileName)

	// 写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	// 实例化
	Logger = logrus.New()

	// 设置输出
	Logger.Out = src

	// 设置日志级别
	Logger.SetLevel(logrus.DebugLevel)

	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		fileName + ".%Y%m%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat:"2006-01-02 15:04:05",
	})

	// 新增 Hook
	Logger.AddHook(lfHook)
}