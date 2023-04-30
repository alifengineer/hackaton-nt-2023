import axios from "axios";

export default axios.create({
  baseURL: `https://${process.env.REACT_APP_API}/admin`,
  headers: {
    Accept: "application/json",
  },
});

export const simpleAxios = axios.create({
  baseURL: `https://${process.env.REACT_APP_API}/admin`,
});

export const authAxios = axios.create({
  baseURL: `https://${process.env.REACT_APP_API}/admin`,
  headers: {
    Accept: "application/json",
  },
});
