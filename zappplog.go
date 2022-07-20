package zappplog

import (
    "io"

    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

var defaultSugarLogger *zap.SugaredLogger
var w io.Writer

func init() {
    zapWs := zapcore.AddSync(&Logger{
        MaxSize:  2048, // megabytes
        Compress: false,
    })

    zapEncoderConfig := zap.NewProductionEncoderConfig()
    zapEncoderConfig.TimeKey = "time"
    zapEncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("20060102-150405")
    zapEncoderConfig.CallerKey = "caller"

    core := zapcore.NewCore(
        zapcore.NewJSONEncoder(zapEncoderConfig),
        zapWs,
        zap.InfoLevel,
    )
    w = zapWs

    defaultSugarLogger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)).Sugar()
}

func Infof(template string, args ...interface{}) {
    defaultSugarLogger.Infof(template, args...)
}

func Errorf(template string, args ...interface{}) {
    defaultSugarLogger.Errorf(template, args...)
}

func Warningf(template string, args ...interface{}) {
    defaultSugarLogger.Warnf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
    defaultSugarLogger.Fatalf(template, args...)
}

func Info(args ...interface{}) {
    defaultSugarLogger.Info(args...)
}

func Error(args ...interface{}) {
    defaultSugarLogger.Error(args...)
}

func Warning(args ...interface{}) {
    defaultSugarLogger.Warn(args...)
}

func Fatal(args ...interface{}) {
    defaultSugarLogger.Fatal(args...)
}

func Infow(msg string, keysAndValues ...interface{}) {
    defaultSugarLogger.Infow(msg, keysAndValues...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
    defaultSugarLogger.Errorw(msg, keysAndValues...)
}

func Warningw(msg string, keysAndValues ...interface{}) {
    defaultSugarLogger.Warnw(msg, keysAndValues...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
    defaultSugarLogger.Fatalw(msg, keysAndValues...)
}

func GetIOWriter() io.Writer {
    return w
}
