package logfile

import (
	"os"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func init() {
	Log.SetReportCaller(true)
	// 设置日志级别为xx以及以上
	Log.SetLevel(logrus.InfoLevel)
	Log.AddHook(&DefaultFieldHook{})
	// 设置日志格式为json格式
	// log.SetFormatter(&logrus.JSONFormatter{
	// 	// PrettyPrint: true,//格式化json
	// 	TimestampFormat: "2006-01-02 15:04:05",//时间格式化
	// })
	Log.SetFormatter(&logrus.TextFormatter{
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		// FullTimestamp:true,
		TimestampFormat: "2006-01-02 15:04:05", //时间格式化
		// DisableLevelTruncation:true,
	})
	// 设置将日志输出到标准输出（默认的输出为stderr，标准错误）
	// 日志消息输出可以是任意的io.writer类型
	Log.SetOutput(os.Stdout)
	//这是我在window上的测试代码，文件名自己修改
	logName := `./logfile`
	writer, err := rotatelogs.New(
		//这是分割代码的命名规则，要和下面WithRotationTime时间精度一致。要是分钟都是分钟
		logName+"%Y%m%d.log",
		// WithLinkName为最新的日志建立软连接，以方便随着找到当前日志文件。windows报错没权限
		// rotatelogs.WithLinkName(logName),

		//文件切割之间的间隔。默认情况下，日志每86400秒/一天旋转一次。注意:记住要利用时间。持续时间值。
		// rotatelogs.WithRotationTime(time.Second*3),
		// WithMaxAge和WithRotationCount二者只能设置一个，
		// WithMaxAge设置文件清理前的最长保存时间，
		// WithRotationCount设置文件清理前最多保存的个数。 默认情况下，此选项是禁用的。
		// rotatelogs.WithMaxAge(time.Second*30),//默认每7天清除下日志文件
		rotatelogs.WithMaxAge(-1),       //需要手动禁用禁用  默认情况下不清除日志，
		rotatelogs.WithRotationCount(2), //清除除最新2个文件之外的日志，默认禁用
	)
	if err != nil {
		Log.Errorf("config local file system for logger error: %v", err)
	}

	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.TextFormatter{DisableColors: true})

	Log.AddHook(lfsHook)
}

type DefaultFieldHook struct {
}

func (hook *DefaultFieldHook) Fire(entry *logrus.Entry) error {
	entry.Data["application"] = "Agent"
	return nil
}

func (hook *DefaultFieldHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

var Loginit = Log.WithFields(logrus.Fields{
	"Server": "Agent",
})
