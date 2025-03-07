package state

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrFailedReadFile = errors.New("Failed to read specified script")
	ErrUnknown        = errors.New("Some very bad error")
)

func GlobalError(err error) {
	fmt.Println("Runtime error:", err)
	os.Exit(-1)
}
