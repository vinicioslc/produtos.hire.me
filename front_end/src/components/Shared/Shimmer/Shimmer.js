import React from "react";
import "./Shimmer.css";
export default class Shimmer extends React.Component {
  state = { show: false };

  render() {
    return (
      <div>
        <div class="shimmer lines"></div>
        <div class="shimmer lines"></div>
        <div class="shimmer lines"></div>
      </div>
    );
  }
}
