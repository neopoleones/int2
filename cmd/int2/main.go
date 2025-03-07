package main

import (
	"context"

	"os/signal"
	"syscall"

	"github.com/destr4ct/int2/internal/int2/state"
	"github.com/destr4ct/int2/pkg/logger"
)

func main() {
	cfg := state.GetConfiguration()
	logger.Setup(cfg.VerboseRun)

	rootContext := context.Background()
	signal.NotifyContext(rootContext, syscall.SIGABRT, syscall.SIGKILL, syscall.SIGTERM)

	if err := Int2Entrypoint(rootContext, &cfg); err != nil {
		state.GlobalError(err)
	}
}
