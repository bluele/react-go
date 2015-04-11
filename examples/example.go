package main

import (
	"fmt"
	"github.com/bluele/react-go"
)

func main() {
	rc, _ := react.NewReact()
	jsx, _ := react.NewJSX()

	component, err := jsx.TransformFile("./example.jsx", map[string]interface{}{
		"harmony":     true,
		"strip_types": true,
	})
	if err != nil {
		panic(err)
	}
	err = rc.Load(component)
	if err != nil {
		panic(err)
	}

	str, err := rc.RenderComponent("App", map[string]interface{}{
		"data": []interface{}{
			map[string]interface{}{
				"id":   1,
				"name": "first",
			},
			map[string]interface{}{
				"id":   2,
				"name": "second",
			},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(str)
}
