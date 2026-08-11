package log

import (
	"io"
	"os"
	"time"

	"github.com/wlbwlbwlb/log/feishu"
	"github.com/wlbwlbwlb/log/logfile"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var writer io.Writer

func Writer() io.Writer {
	return writer
}

var logger *zap.SugaredLogger

func Logger() *zap.SugaredLogger {
	return logger
}

var opt Options

func Init(name string, opts ...OptionFunc) (logger *zap.SugaredLogger, e error) {
	opt = Options{
		Name: name,
	}
	for _, fn := range opts {
		fn(&opt)
	}

	encoder := zapcore.NewJSONEncoder(newEncoderConfig())

	writer = io.MultiWriter(os.Stdout, logfile.Writer)

	logger = zap.New(
		zapcore.NewTee(
			zapcore.NewCore(encoder,
				zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(logfile.Writer)),
				zapcore.InfoLevel,
			),
			zapcore.NewCore(encoder, zapcore.AddSync(feishu.Writer), zapcore.ErrorLevel),
		),
		zap.Fields(zap.String("name", opt.Name)),
		zap.AddStacktrace(zapcore.ErrorLevel),
		zap.AddCallerSkip(1),
		zap.AddCaller(),
	).Sugar()

	return
}

func newEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout(time.DateTime),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func Info(args ...interface{}) {
	if nil == logger {
		panic("init first")
	}
	logger.Info(args...)
}

func Warn(args ...interface{}) {
	if nil == logger {
		panic("init first")
	}
	logger.Warn(args...)
}

func Error(args ...interface{}) {
	if nil == logger {
		panic("init first")
	}
	logger.Error(args...)
}

func Panic(args ...interface{}) {
	if nil == logger {
		panic("init first")
	}
	logger.Panic(args...)
}

func Infof(template string, args ...interface{}) {
	if nil == logger {
		panic("init first")
	}
	logger.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	if nil == logger {
		panic("init first")
	}
	logger.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	if nil == logger {
		panic("init first")
	}
	logger.Errorf(template, args...)
}

func Panicf(template string, args ...interface{}) {
	if nil == logger {
		panic("init first")
	}
	logger.Panicf(template, args...)
}

func With(args ...interface{}) *Wrap {
	if nil == logger {
		panic("init first")
	}
	return &Wrap{
		SugaredLogger: logger.With(args...),
	}
}
