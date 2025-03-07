package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/destr4ct/int2/internal/int2/ast/T"
	"github.com/destr4ct/int2/internal/int2/scanner"
	"github.com/destr4ct/int2/internal/int2/state"
	"github.com/destr4ct/int2/internal/int2/token"
)

func getTokens(raw string, cfg *state.Int2Configuration) ([]*token.Token, error) {
	tokens := cfg.Scanner.Tokenize(raw)
	if err := scanner.Validate(tokens); err != nil {
		return nil, err
	}

	if cfg.VerboseRun {
		token.PrintTokens(tokens)
	}

	return tokens, nil
}

func getASTReproExpr(raw string, cfg *state.Int2Configuration) (T.Expr, error) {
	tokens, err := getTokens(raw, cfg)
	if err != nil {
		return nil, err
	}

	// Parse to AST
	expr, err := cfg.Parser.ParseExpr(tokens)
	if err != nil {
		return nil, err
	}

	return expr, nil
}

func getASTRepro(raw string, cfg *state.Int2Configuration) ([]T.Stmt, error) {
	tokens, err := getTokens(raw, cfg)
	if err != nil {
		return nil, err
	}

	return cfg.Parser.Parse(tokens)
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")

	inp, _ := reader.ReadString('\n')
	return strings.TrimRight(inp, "\n")
}
