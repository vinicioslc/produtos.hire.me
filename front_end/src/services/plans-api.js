import axios from "axios";

export default {
  getPlans: async function(carrier) {
    return (await axios.get(
      `http://localhost:8080/v1/carrier/${carrier}/plans`
    )).data;
  }
};
