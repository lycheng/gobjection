package signals

import (
	"os"
	"os/signal"
	"syscall"
)

// WaitToStopped waits for the INT TERM signals
func WaitToStopped() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
}
