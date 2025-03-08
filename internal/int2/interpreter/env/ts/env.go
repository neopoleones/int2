/*
	Thread-safe storage of env variables (maybe I will implement multithreadng for int2 in future?)
	TODO: get the way I need to work with environCtx
*/

package ts

import (
	"fmt"
	"sync"

	"github.com/destr4ct/int2/internal/int2/interpreter"
)

type Int2Environ struct {
	mut     *sync.RWMutex
	storage map[string]any
}

func (e *Int2Environ) Get(k string, _ *interpreter.EnvironCtx) (any, error) {
	e.mut.RLock()
	defer e.mut.RUnlock()

	v, ok := e.storage[k]
	if !ok {
		return nil, &interpreter.RuntimeError{
			Reason: fmt.Sprintf("unknown variable: %s", k),
		}
	}

	return v, nil
}

func (e *Int2Environ) Set(k string, v any, _ *interpreter.EnvironCtx) error {
	e.mut.Lock()
	defer e.mut.Unlock()

	e.storage[k] = v
	return nil
}

func NewEnv() *Int2Environ {
	return &Int2Environ{
		mut:     new(sync.RWMutex),
		storage: make(map[string]any),
	}
}
