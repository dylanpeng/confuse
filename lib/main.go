package main

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"time"
)

func main() {
	//conf := zap.NewDevelopmentConfig()
	//conf.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	//conf.OutputPaths = []string{"logs/log.log", "stderr"}
	//
	//p, _ := conf.Build()
	//defer p.Sync()
	//ps := p.Sugar()
	//
	//p.WithOptions()
	//
	//p.Info("NewDevelopment", zap.Int("int", 10))
	//ps.Info("Dev sugar", "test")
	//ps.Infof("dev sugar f %s", "aaa")
	//ps.Infof("dev sugar d")
	//ps.Infow("dev sugar w ", "url", "baidu")
	//ps.Error("error")

	encoderConf := zap.NewDevelopmentEncoderConfig()
	encoderConf.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	rotate := getWriter("./logs/confuse")
	writer := zapcore.AddSync(rotate)

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConf),
		writer,
		zap.DebugLevel,
	)

	p := zap.New(core, zap.AddStacktrace(zap.WarnLevel), zap.AddCaller())
	ps := p.Sugar()

	ps.Info("Dev sugar", "test")
	ps.Infoln("Dev sugar", "test")
	ps.Infof("dev sugar f %s", "aaa")
	//ps.Infof("dev sugar d")
	//ps.Infow("dev sugar w ", "url", "baidu")
	//ps.Error("error")

}

func getWriter(filename string) io.Writer {
	// 生成rotatelogs的Logger 实际生成的文件名 demo.log.YYmmddHH
	// demo.log是指向最新日志的链接
	// 保存7天内的日志，每1小时(整点)分割一次日志
	hook, err := rotatelogs.New(
		// 没有使用go风格反人类的format格式
		filename+".%Y-%m-%d.log",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*30),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		panic(err)
	}
	return hook
}
