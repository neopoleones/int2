package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/destr4ct/int2/internal/int2/interpreter"
	"github.com/destr4ct/int2/internal/int2/state"
	"github.com/destr4ct/int2/internal/int2/utils"
	"github.com/destr4ct/int2/pkg/logger"
)

const (
	TOKEN_EXIT = ".exit"
)

func runtimeErrorHandler() {
	raw := recover()

	if err, ok := raw.(*interpreter.RuntimeError); ok {
		state.GlobalError(err)
	}

	if raw != nil {
		panic(raw)
	}
}

func RunInterpreter(ctx context.Context, cfg *state.Int2Configuration) error {
	defer runtimeErrorHandler()

	for inp := getUserInput(); inp != TOKEN_EXIT; inp = getUserInput() {
		ast, err := getASTRepresentationOfExpr(inp, cfg)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("<", cfg.Interpreter.Evaluate(ast))
	}

	return nil
}

func RunScript(ctx context.Context, cfg *state.Int2Configuration) error {
	defer runtimeErrorHandler()

	scriptContent, err := utils.ReadFile(cfg.ScriptPath)
	if err != nil {
		return state.ErrFailedReadFile
	}

	ast, err := getASTRepresentationOfExpr(scriptContent, cfg)
	if err != nil {
		return err
	}

	res := cfg.Interpreter.Evaluate(ast)
	fmt.Println("res:", res)

	return nil
}

func Int2Entrypoint(ctx context.Context, cfg *state.Int2Configuration) error {
	logger.Base.Debug("Starting the interpreter", slog.Attr{
		Key:   "options",
		Value: slog.AnyValue(*cfg),
	})

	if len(cfg.ScriptPath) == 0 {
		return RunInterpreter(ctx, cfg)
	}

	return RunScript(ctx, cfg)
}
