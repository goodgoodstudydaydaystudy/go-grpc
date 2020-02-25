package sig

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func SignalDone(msg string) {
	s := make(chan os.Signal)
	signal.Notify(s, syscall.SIGINT, syscall.SIGCONT)

	go func() {
		for {
			_ = <- s
			log.Println(msg)
		}
	}()
}