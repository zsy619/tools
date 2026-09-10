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

// LoggerPath 默认日志根目录，可在 init 之后修改。
var LoggerPath string

func init() {
	LoggerPath = "./logs"
}

// Logger 通用日志对象，包装 *logrus.Logger 以便直接使用其全部方法。
type Logger struct {
	*logrus.Logger
}

// RouteLogger 通道日志，按通道类型分别记录不同动作的日志。
type RouteLogger struct {
	Get        *logrus.Logger // 提取
	Submit     *logrus.Logger // 提交
	SubmitResp *logrus.Logger // 提交回应
	ActiveTest *logrus.Logger // 心跳包
	ReportMO   *logrus.Logger // 状态上行
	Deliver    *logrus.Logger // 回应
	Info       *logrus.Logger // 其他信息
}

// NewRouteLogger 为指定通道（routeCode + linkIndex）构造一组按动作分类的通道日志。
// 每类日志会写到 LoggerPath/<routeCode>/<linkIndex>/<action> 目录下。
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

// NewLogger 在 LoggerPath/basePath 下构造通用 Logger。
// 目录创建失败会打印到标准输出，但不会中断 Logger 构造。
func NewLogger(basePath string) *Logger {
	baseLogPath := fmt.Sprintf("%s/%s", LoggerPath, basePath)
	err := os.MkdirAll(baseLogPath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	return &Logger{newLogger(baseLogPath)}
}

// newRouteLogger 在 LoggerPath/<routeCode>/<linkIndex>/<basePath> 下创建分类日志。
func newRouteLogger(routeCode, basePath string, linkIndex int) *logrus.Logger {
	baseLogPath := fmt.Sprintf("%s/%s/%d/%s", LoggerPath, routeCode, linkIndex, basePath)
	err := os.MkdirAll(baseLogPath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	return newLogger(baseLogPath)
}

// newLogger 构造一个 Info 级别、按天切分的 logrus.Logger。
// 默认保留 7 天，单个文件最大 1 天后切分。
func newLogger(baseLogPath string) *logrus.Logger {
	loginfo := logrus.New()
	loginfo.SetLevel(logrus.InfoLevel)
	loginfo.AddHook(newRotateHook(baseLogPath, 7*24*time.Hour, 24*time.Hour))
	return loginfo
}

// newRotateHook 为 debug/info/warn/error/fatal/panic 各级别分别创建滚动日志 writer。
// maxAge 为日志最大保留时间，rotationTime 为切分间隔。
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

// newLogLevel 为单个日志级别创建 rotatelogs.RotateLogs 写入器。
// 文件名格式：<logPath>/<logFileName>.%Y%m%d.log，并通过软链指向最新日志文件。
// 创建失败时打印错误日志到标准输出，但仍会返回（可能为 nil）writer。
func newLogLevel(logPath, logFileName string, maxAge time.Duration, rotationTime time.Duration) *rotatelogs.RotateLogs {
	baseLogPath := path.Join(logPath, logFileName)
	writer, err := rotatelogs.New(
		baseLogPath+".%Y%m%d.log",                 // 日志文件格式
		rotatelogs.WithLinkName(baseLogPath),      // 生成软链，指向最新日志文件
		rotatelogs.WithRotationTime(time.Hour*24), // 设置日志分割时间（此处设为 1 小时）
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	if err != nil {
		errx := fmt.Errorf("config local file system logger error. %+v", errors.WithStack(err))
		fmt.Println(errx)
	}
	return writer
}
