package log

import (
	"io"
)

type Logger interface {
	Configure(conf *LogConf) error
}

type LogLevel int;
const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

type LogConf struct {
	Level LogLevel
	Format string
	Output io.Writer
	LogPath string
	LogFile string
	MaxSizePerFile int
	MaxBackups int
	Compress bool
}

func NewLogConf() *LogConf {
	return &LogConf{
		Level: InfoLevel,
		Format: "",
		Output: nil,
		LogPath: "./",
		LogFile: "app.log",
		MaxSizePerFile: 100,
		MaxBackups: 10,
		Compress: true,
	}
}