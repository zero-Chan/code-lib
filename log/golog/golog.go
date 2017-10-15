package golog

import (
	"io"
	"log"
	"os"

	slog "code-lib/log"
)

type GoLog struct {
	*log.Logger
}

func Create(out io.Writer, prefix string, flag int) GoLog {
	golog := GoLog{
		Logger: log.New(out, prefix, flag),
	}

	return golog
}

func New(out io.Writer, prefix string, flag int) *GoLog {
	golog := Create(out, prefix, flag)
	return &golog
}

func (this *GoLog) Virtualize() slog.Logger {
	return this
}

func CreateDefault() GoLog {
	return Create(os.Stderr, "", log.LstdFlags)
}

func NewDefault() *GoLog {
	defaultLog := CreateDefault()
	return &defaultLog
}

func (this *GoLog) Error(v ...interface{}) {
	this.Logger.Fatal(v...)
}

func (this *GoLog) Errorf(format string, v ...interface{}) {
	this.Logger.Fatalf(format, v...)
}

func (this *GoLog) Errorln(v ...interface{}) {
	this.Logger.Fatalln(v...)
}

func (this *GoLog) Notice(v ...interface{}) {
	this.Logger.Print(v...)
}

func (this *GoLog) Noticef(format string, v ...interface{}) {
	this.Logger.Printf(format, v...)
}

func (this *GoLog) Noticeln(v ...interface{}) {
	this.Logger.Println(v...)
}
