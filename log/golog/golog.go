package golog

import (
	"io"
	"log"
	"os"

	slog "subassembly/log"
)

type GoLog struct {
	log *log.Logger
}

func Create(out io.Writer, prefix string, flag int) GoLog {
	golog := GoLog{
		log: log.New(out, prefix, flag),
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

func (this *GoLog) Printf(format string, v ...interface{}) {
	this.log.Printf(format, v...)
}

func (this *GoLog) Print(v ...interface{}) {
	this.log.Print(v...)
}

func (this *GoLog) Println(v ...interface{}) {
	this.log.Println(v...)
}
