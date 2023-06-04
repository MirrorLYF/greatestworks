package signal

import (
	"greatestworks/aop/logger"
	"os"
	"os/signal"
	"syscall"
)

func OnStart(o Owner) {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)
	//go func() {
	for s := range sig {
		switch s {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL:
			logger.InfoF("received signal :%v", s)
			o.Stop()
		case syscall.SIGABRT:
			logger.InfoF("[shutdown command ] received signal :%v", s)
			o.Stop()
		default:
			logger.InfoF("received signal :%v", s)
		}
	}
	//}()
}
