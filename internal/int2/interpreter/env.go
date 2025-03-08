package interpreter

var (
	EnvBackgroundCtx = NewEnvCtx()
)

type EnvironCtx struct {
}

func NewEnvCtx() *EnvironCtx {
	return &EnvironCtx{}
}

type Environ interface {
	Get(string, *EnvironCtx) (any, error)
	Set(string, any, *EnvironCtx) error
}
