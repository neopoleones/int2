package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/destr4ct/int2/internal/int2/ast/visitor"
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
		ast, err := getASTReproExpr(inp, cfg)
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

	astStmt, err := getASTRepro(scriptContent, cfg)
	if err != nil {
		return err
	}

	if cfg.VerboseRun {
		fmt.Println("AST:")
		for i, stmt := range astStmt {
			fmt.Println(i, visitor.NewAstPrinter().Stringify(stmt))
		}
	}

	cfg.Interpreter.Interpret(astStmt)
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
