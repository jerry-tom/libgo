package log

import (
	"path/filepath"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	logger *logrus.Logger
}

func NewLogrusLogger() *LogrusLogger {
	l := &LogrusLogger{
		logger: logrus.New(),
	}
	conf := NewLogConf()

	l.Configure(conf)
	return l
}

func NewLogrusLoggerWithLevel(level LogLevel) *LogrusLogger {
	l := NewLogrusLogger()
	l.setLevel(level)

	return l
}

func (l *LogrusLogger) Configure(conf *LogConf) error {
	// set log level
	l.setLevel(conf.Level)

	// set log format
	formatter := new(logrus.TextFormatter)
	formatter.FullTimestamp = true
	formatter.Caller = true
	l.logger.SetFormatter(formatter)

	// set log rollback
	if conf.LogPath == "" {
		conf.LogPath = "./"
	}

	if conf.LogFile == "" {
		conf.LogFile = "app.log"
	}

	logFile := &lumberjack.Logger{
		Filename:   filepath.Join(conf.LogPath, conf.LogFile),
		MaxSize:    conf.MaxSizePerFile,
		MaxBackups: conf.MaxBackups,
		MaxAge:     30,
		Compress:   conf.Compress,
	}
	l.logger.AddHook(logrus.NewFileHook(logFile))

	return nil
}

func (l *LogrusLogger) setLevel(level LogLevel) {
	switch level {
	case DebugLevel:
		l.logger.Level = logrus.DebugLevel
	case InfoLevel:
		l.logger.Level = logrus.InfoLevel
	case WarnLevel:
		l.logger.Level = logrus.WarnLevel
	case ErrorLevel:
		l.logger.Level = logrus.ErrorLevel
	default:
		l.logger.Level = logrus.InfoLevel
	}
}
