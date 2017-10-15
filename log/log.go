package log

type Logger interface {
	Virtualize() Logger

	// print: 无tag输出
	Printf(format string, v ...interface{})
	Print(v ...interface{})
	Println(v ...interface{})

	// TODO
	//	// fatal: 严重的错误导致程序必须关闭的错误
	//	Fatal(v ...interface{})
	//	Fatalf(format string, v ...interface{})
	//	Fatalln(v ...interface{})

	//	// panic: 启动时程序初始化失败导致程序退出的错误
	//	Panic(v ...interface{})
	//	Panicf(format string, v ...interface{})
	//	Panicln(v ...interface{})

	//	// crit: 紧急的错误,程序性运行的错误,需要人工马上参与修复的错误
	//	Crit(v ...interface{})
	//	Critf(format string, v ...interface{})
	//	Critln(v ...interface{})

	//	// error: 用户传入参数错误等一般等级的错误消息
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
	Errorln(v ...interface{})

	//	// warn: 出错但不影响程序继续运行,程序能自己恢复运行,不需要人工参与修复的问题
	//	Warn(v ...interface{})
	//	Warnf(format string, v ...interface{})
	//	Warnln(v ...interface{})

	//	// notice:  接口访问消息,用于统计访问量,需要设计好足够的统计参数
	Notice(v ...interface{})
	Noticef(format string, v ...interface{})
	Noticeln(v ...interface{})

	//	// info: 用户传入的参数,运行状态输出等辅助跟踪错误的信息
	//	Info(v ...interface{})
	//	Infof(format string, v ...interface{})
	//	Infoln(v ...interface{})

	//	// debug: 仅用于调试
	//	Debug(v ...interface{})
	//	Debugf(format string, v ...interface{})
	//	Debugln(v ...interface{})
}
