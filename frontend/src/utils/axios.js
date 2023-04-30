import axios from "axios";

export default axios.create({
  baseURL: `http://${process.env.REACT_APP_API}`,
  headers: {
    Accept: "application/json",
  },
});
