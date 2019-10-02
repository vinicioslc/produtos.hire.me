import React from "react";

/*
 * MainPage serves the purpose of a "root route component".
 */
class MainPage extends React.Component {
  /**
   * Here, we define a react lifecycle method that gets executed each time
   * our component is mounted to the DOM, which is exactly what we want in this case
   */
  componentDidMount() {
    document.title = this.props.title;
  }

  /**
   * Here, we use a component prop to render
   * a component, as specified in route configuration
   */
  render() {
    const PageComponent = this.props.component;

    return <PageComponent />;
  }
}

export default MainPage;
