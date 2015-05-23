# react-go

react-go is a Go wrapper around the [React](http://facebook.github.io/react/) and [JSX](http://facebook.github.io/react/docs/jsx-in-depth.html). It enables you to do server-side rendering with the React.

## Overview

Currently react.js and JSXTransformer version are *0.13.3*.

*Dependencies:* react-go use [go-duktape](https://github.com/olebedev/go-duktape) to evaluate javascript.

## Features

* Support React serverside rendering.
* Transform JSX to JS.

## Example

```go
func main() {
  rc, _ := react.NewReact()
  jsx, _ := react.NewJSX()

  component, _ := jsx.TransformFile("./example.jsx", map[string]interface{}{
    "harmony":     true,
    "strip_types": true,
  })
  rc.Load(component)

  str, _ := rc.RenderComponent("App", map[string]interface{}{
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
  fmt.Println(str)
}
```

```jsx
'use strict';
var React = self.React;

var App = React.createClass({
  getInitialState() {
    return {
      message: "loading..."
    };
  },
  componentDidMount() {
    this.setState({ message: "welcome!" });
  },
  render() {
    var list = this.props.data.map(obj => <li key={obj.id}>{obj.name}</li>);
    return (
      <div>
        <p>server-side rendering sample</p>
        <p>{this.state.message}</p>
        <ul>{list}</ul>
      </div>
    );
  }
});

self.App = App;
```

output (formatted):
```html
<div data-reactid=".1fei395a42" data-react-checksum="-1544136830">
  <p data-reactid=".1fei395a42.0">server-side rendering sample</p>
  <p data-reactid=".1fei395a42.1">loading...</p>
  <ul data-reactid=".1fei395a42.2">
    <li data-reactid=".1fei395a42.2.$1">first</li>
    <li data-reactid=".1fei395a42.2.$2">second</li>
  </ul>
</div>
```

## Getting Started

### Install

```
$ go get -u github.com/bluele/react-go
```

# Author

**Jun Kimura**

* <http://github.com/bluele>
* <junkxdev@gmail.com>