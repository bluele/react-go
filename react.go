package react

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type React struct {
	pool *Pool
	opt  *Option
}

// Create a new react object.
func NewReact() (*React, error) {
	return NewReactWithOption(DefaultReactOption())
}

// Create a new react object using option.
// opt: Option for react object.
func NewReactWithOption(opt *Option) (*React, error) {
	if opt == nil {
		return nil, errors.New("react: nil opt *Option")
	}
	err := opt.Validate()
	if err != nil {
		return nil, err
	}

	pool, err := NewPool(opt)
	if err != nil {
		return nil, err
	}

	return &React{pool: pool, opt: opt}, nil
}

// Returns a default option for react.
func DefaultReactOption() *Option {
	src, err := Asset("assets/react.js")
	if err != nil {
		panic(err)
	}
	return &Option{
		Source:           src,
		PoolSize:         10,
		GlobalObjectName: "self",
	}
}

// Render react component.
// name: component name
// params: component properties
func (rc *React) RenderComponent(name string, params interface{}) (string, error) {
	vm := rc.pool.Get()
	defer rc.pool.Put(vm)

	objName := rc.opt.GlobalObjectName

	var js string
	if params == nil {
		js = fmt.Sprintf(`
			%v.React.renderToString(
				%v.React.createFactory(%v.%v)()
			)`, objName, objName, objName, name)
	} else {
		j, err := json.Marshal(params)
		if err != nil {
			return "", err
		}
		js = fmt.Sprintf(`
			%v.React.renderToString(
				%v.React.createFactory(%v.%v)(%v)
			)`, objName, objName, objName, name, string(j))
	}

	if vm.PevalString(js) == 1 {
		return "", errors.New(vm.SafeToString(-1))
	}
	v := vm.SafeToString(-1)
	vm.Pop()
	return v, nil
}

// Run javascript code and returns its result value.
// src: javascript source
func (rc *React) RunScript(src string) (string, error) {
	vm := rc.pool.Get()
	defer rc.pool.Put(vm)

	js := `
	(function() {
		return %v
	})();
	`
	if vm.PevalString(fmt.Sprintf(js, src)) == 1 {
		return "", errors.New(vm.SafeToString(-1))
	}
	ret := vm.SafeToString(-1)
	vm.Pop()
	return ret, nil
}

// Load javascript code.
// src: javascript source
func (rc *React) Load(src []byte) error {
	for i := 0; i < rc.pool.size; i++ {
		vm := rc.pool.Get()
		defer rc.pool.Put(vm)

		if vm.PevalString(string(src)) == 1 {
			panic(vm.SafeToString(-1))
		}
	}
	return nil
}

// Load javascript file.
// path: path for javascript source file
func (rc *React) LoadFile(path string) error {
	src, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return rc.Load(src)
}
