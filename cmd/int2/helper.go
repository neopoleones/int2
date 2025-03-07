package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/destr4ct/int2/internal/int2/ast/T"
	"github.com/destr4ct/int2/internal/int2/ast/visitor"
	"github.com/destr4ct/int2/internal/int2/scanner"
	"github.com/destr4ct/int2/internal/int2/state"
	"github.com/destr4ct/int2/internal/int2/token"
)

func getASTRepresentationOfExpr(raw string, cfg *state.Int2Configuration) (T.Expr, error) {
	tokens := cfg.Scanner.Tokenize(raw)
	if err := scanner.Validate(tokens); err != nil {
		return nil, err
	}

	if cfg.VerboseRun {
		token.PrintTokens(tokens)
	}

	// Parse to AST
	expr, err := cfg.Parser.ParseExpr(tokens)
	if err != nil {
		return nil, err
	}

	if cfg.VerboseRun {
		fmt.Println("AST:", visitor.NewAstPrinter().Stringify(expr))
	}

	return expr, nil
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")

	inp, _ := reader.ReadString('\n')
	return strings.TrimRight(inp, "\n")
}
