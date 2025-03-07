package state

import (
	"flag"
	"sync"

	"github.com/destr4ct/int2/internal/int2/ast"
	"github.com/destr4ct/int2/internal/int2/ast/rdparser"
	"github.com/destr4ct/int2/internal/int2/interpreter"
	"github.com/destr4ct/int2/internal/int2/interpreter/evaluator"
	"github.com/destr4ct/int2/internal/int2/scanner"
	"github.com/destr4ct/int2/internal/int2/scanner/baseline"
)

var once sync.Once
var cfg Int2Configuration

var chooseList = []func(){
	chooseParser,
	chooseScanner,
	chooseInterpreter,
}

func init() {
	flag.BoolVar(&cfg.VerboseRun, "verbose", false, "run interpreter in verbose mode")
	flag.StringVar(&cfg.scannerType, "scanner", "baseline", "scanner used to tokenize script")
	flag.StringVar(&cfg.parserType, "parser", "rd", "parser used to create AST from tokens")
}

type Int2Configuration struct {
	VerboseRun bool

	ScriptPath string

	scannerType string
	Scanner     scanner.Scanner

	parserType string
	Parser     ast.Parser

	Interpreter interpreter.Interpreter
}

func chooseScanner() {
	switch cfg.scannerType {
	default:
		cfg.Scanner = baseline.GetScanner()
	}
}

func chooseParser() {
	switch cfg.parserType {
	default:
		cfg.Parser = rdparser.Get()
	}
}

func chooseInterpreter() {
	cfg.Interpreter = evaluator.Get()
}

func GetConfiguration() Int2Configuration {
	once.Do(func() {
		flag.Parse()

		if len(flag.Args()) > 0 {
			cfg.ScriptPath = flag.Arg(0)
		}

		for _, choose := range chooseList {
			choose()
		}
	})

	return cfg
}
