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