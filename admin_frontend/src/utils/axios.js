import axios from "axios";

export default axios.create({
  baseURL: `https://${process.env.REACT_APP_API}`,
  headers: {
    Accept: "application/json",
  },
});

export const simpleAxios = axios.create({
  baseURL: `https://${process.env.REACT_APP_API}`,
});

export const authAxios = axios.create({
  baseURL: `https://${process.env.API}`,
  headers: {
    Accept: "application/json",
  },
});
