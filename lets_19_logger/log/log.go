package log

import (
	"os"
)

type Logger interface {
	Info(msg string)
	Error(msg string)
	DeBug(msg string)
}

//file文件输出log
type Flog struct {
	FilePath string //log file path
}

func (f Flog) Info(msg string) {
	OpenWriter(f).WriteString("[INFO]:" + msg + "\n")
}
func (f Flog) Error(msg string) {
	OpenWriter(f).WriteString("[ERROR]:" + msg + "\n")
}

func (f Flog) DeBug(msg string) {
	OpenWriter(f).WriteString("[DeBug]:" + msg + "\n")
}

//console控制台log
type Clog struct {
}

func (c Clog) Info(msg string) {
	os.Stdout.WriteString("[INFO]:" + msg + "\n")
}
func (c Clog) Error(msg string) {
	os.Stdout.WriteString("[Error]:" + msg + "\n")
}
func (c Clog) DeBug(msg string) {
	os.Stdout.WriteString("[INFO]:" + msg + "\n")
}

func NewFlog(file string) Logger {
	f := &Flog{file}
	logger := Logger(f)
	return logger
}
func NewClog() Logger {
	logger := Logger(&Clog{})
	return logger
}
func OpenWriter(f Flog) *os.File {
	file, err := os.OpenFile(f.FilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	defer file.Close()
	if err != nil {
		panic("日志输出错误...")
	}
	return file
}
