package react

import (
	"errors"
	"fmt"
	"github.com/olebedev/go-duktape"
)

type Pool struct {
	size int
	ch   chan *duktape.Context
}

type Option struct {
	Source []byte
	// size for javascript vm pool.
	PoolSize int
	// name for variable includes component objects. ex. "self"
	GlobalObjectName string
}

func (opt *Option) Validate() error {
	if opt.Source == nil {
		return errors.New("react: nil []byte opt.Source")
	}
	if opt.PoolSize <= 0 {
		return errors.New("react: opt.PoolSize must be greater than or equal to 1")
	}
	if opt.GlobalObjectName == "" {
		return errors.New("react: empty string opt.GlobalObjectName")
	}
	return nil
}

func NewPool(opt *Option) (*Pool, error) {
	pool := &Pool{size: opt.PoolSize}
	pool.ch = make(chan *duktape.Context, opt.PoolSize)
	for i := 0; i < pool.size; i++ {
		vm, err := newVM(opt.Source, opt.GlobalObjectName)
		if err != nil {
			return nil, err
		}
		pool.ch <- vm
	}
	return pool, nil
}

func newVM(src []byte, objName string) (*duktape.Context, error) {
	vm := duktape.NewContext()
	source := fmt.Sprintf(`var %v = %v || {}, console = {log:print,warn:print,error:print,info:print}; `, objName, objName) + string(src)
	if vm.PevalString(source) == 1 {
		return nil, errors.New(vm.SafeToString(-1))
	}
	return vm, nil
}

func (pl *Pool) Get() *duktape.Context {
	return <-pl.ch
}

func (pl *Pool) Put(vm *duktape.Context) {
	pl.ch <- vm
}
