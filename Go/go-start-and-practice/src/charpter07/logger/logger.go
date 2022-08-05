package main

import (
	"errors"
	"fmt"
	"os"
)

// 声明日志显示接口
type LogWriter interface {
	Write(data interface{}) error
}

// 日志器
type Logger struct {
	// 日志器用到了日志写入器
	writerList []LogWriter
}

// 注册一个日志写入设备
func (l *Logger) RegisterWriter(writer LogWriter) {
	l.writerList = append(l.writerList, writer)
}

// 将一个 data 类型的数据写入日志
func (l *Logger) Log(data interface{}) {
	// 遍历所有注册的写入器
	for _, writer := range l.writerList {
		writer.Write(data)
	}
}

// 创建日志器的实例
func NewLogger() *Logger {
	return &Logger{}
}

// 文件写入器是众多日志写入器的一种，其他的比如写入到标准输出

// 声明文件写入器
type fileWriter struct {
	file *os.File
}

func (f *fileWriter) SetFile(filename string) (err error) {
	// 文件以及打开，关闭前一个文件
	if f.file != nil {
		f.file.Close()
	}
	f.file, err = os.Create(filename)
	return err
}

// 实现接口的方法
func (f *fileWriter) Write(data interface{}) error {
	if f.file == nil {
		return errors.New("file not created")
	}

	// 按照 data 本来的值转换为字符串
	str := fmt.Sprintf("%v\n", data)
	// 将 str 字符串转换为 []byte 数组
	_, err := f.file.Write([]byte(str))
	return err
}

// 创建写入文件器的实例
func newFileWriter() *fileWriter {
	return &fileWriter{}
}

// 命令行写入器，go 语言中标准输入输出，标准错误输出都是文件 *os.File 类型
type consoleWriter struct {
}

// 实现接口的 Write 方法
func (f *consoleWriter) Write(data interface{}) error {
	str := fmt.Sprintf("%v\n", data)
	_, err := os.Stdout.Write([]byte(str))
	return err
}

func newConsoleWriter() *consoleWriter {
	return &consoleWriter{}
}

// 创建日志器
func createLogger() *Logger {
	// 创建日志器
	l := NewLogger()
	// 创建命令行写入器
	cw := newConsoleWriter()
	// 注册到日志写入器中
	l.RegisterWriter(cw)
	// 创建文件写入器
	fw := newFileWriter()
	// 设置文件名
	if err := fw.SetFile("log.log"); err != nil {
		fmt.Println(err)
	}
	// 注册到日志写入器中
	l.RegisterWriter(fw)
	return l
}

func main() {
	l := createLogger()
	l.Log("hello")
}
