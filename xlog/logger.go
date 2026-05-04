package xlog

import (
	"fmt"
	"os"
	"path"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/pkg/errors"
	lfshook "github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

var LoggerPath string

func init() {
	LoggerPath = "./logs"
}

// Logger 通用日志信息
type Logger struct {
	*logrus.Logger
}

// RouteLogger 通道日志
type RouteLogger struct {
	Get        *logrus.Logger // 提取
	Submit     *logrus.Logger // 提交
	SubmitResp *logrus.Logger // 提交回应
	ActiveTest *logrus.Logger // 心跳包
	ReportMO   *logrus.Logger // 状态上行
	Deliver    *logrus.Logger // 回应
	Info       *logrus.Logger // 其他信息
}

// NewRouteLogger 构造通道日志
func NewRouteLogger(routeCode string, linkIndex int) *RouteLogger {
	result := new(RouteLogger)
	result.Info = newRouteLogger(routeCode, "info", linkIndex)
	result.Get = newRouteLogger(routeCode, "get", linkIndex)
	result.Submit = newRouteLogger(routeCode, "submit", linkIndex)
	result.SubmitResp = newRouteLogger(routeCode, "submit_resp", linkIndex)
	result.ActiveTest = newRouteLogger(routeCode, "active", linkIndex)
	result.ReportMO = newRouteLogger(routeCode, "report_mo", linkIndex)
	result.Deliver = newRouteLogger(routeCode, "deliver", linkIndex)
	return result
}

// NewLogger 构造日志
func NewLogger(basePath string) *Logger {
	baseLogPath := fmt.Sprintf("%s/%s", LoggerPath, basePath)
	err := os.MkdirAll(baseLogPath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	return &Logger{newLogger(baseLogPath)}
}

func newRouteLogger(routeCode, basePath string, linkIndex int) *logrus.Logger {
	baseLogPath := fmt.Sprintf("%s/%s/%d/%s", LoggerPath, routeCode, linkIndex, basePath)
	err := os.MkdirAll(baseLogPath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	return newLogger(baseLogPath)
}

func newLogger(baseLogPath string) *logrus.Logger {
	loginfo := logrus.New()
	loginfo.SetLevel(logrus.InfoLevel)
	loginfo.AddHook(newRotateHook(baseLogPath, 7*24*time.Hour, 24*time.Hour))
	return loginfo
}

func newRotateHook(logPath string, maxAge, rotationTime time.Duration) *lfshook.LfsHook {
	d := newLogLevel(logPath, "debug", maxAge, rotationTime) // 调试
	i := newLogLevel(logPath, "info", maxAge, rotationTime)  // 信息
	w := newLogLevel(logPath, "warn", maxAge, rotationTime)  // 警告
	e := newLogLevel(logPath, "error", maxAge, rotationTime) // 错误
	f := newLogLevel(logPath, "fetal", maxAge, rotationTime) // 异常
	p := newLogLevel(logPath, "panic", maxAge, rotationTime) // panic

	return lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: d,
		logrus.InfoLevel:  i,
		logrus.WarnLevel:  w,
		logrus.ErrorLevel: e,
		logrus.FatalLevel: f,
		logrus.PanicLevel: p,
	},
		&easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05.000", // 时间格式
			LogFormat:       "%time% - %msg%\r",        // 消息格式
		})
}

func newLogLevel(logPath, logFileName string, maxAge time.Duration, rotationTime time.Duration) *rotatelogs.RotateLogs {
	baseLogPath := path.Join(logPath, logFileName)
	writer, err := rotatelogs.New(
		baseLogPath+".%Y%m%d.log",                 // 日志文件格式
		rotatelogs.WithLinkName(baseLogPath),      // 生成软链，指向最新日志文
		rotatelogs.WithRotationTime(time.Hour*24), // WithRotationTime 設置日誌分割的時間，這裏設置爲一小時分割一次
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	if err != nil {
		errx := fmt.Errorf("config local file system logger error. %+v", errors.WithStack(err))
		fmt.Println(errx)
	}
	return writer
}
