package logger

import (
	"log/slog"
	"os"
)

var Base *slog.Logger

func Setup(verbose bool) {
	var hdl slog.Handler
	
	if(verbose) {
		hdl = slog.NewTextHandler(
			os.Stdout, &slog.HandlerOptions{
				AddSource: true,
				Level: slog.LevelDebug,
			},
		)
	} else {	
		hdl = slog.NewTextHandler(
			os.Stderr, &slog.HandlerOptions{
				Level: slog.LevelError,
			},
		)
	}
	
	Base = slog.New(hdl)
}