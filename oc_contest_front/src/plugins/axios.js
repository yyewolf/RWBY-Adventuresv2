import axios from "axios";

let base_backend_url = `${process.env.VUE_APP_BACKEND}`;
axios.defaults.withCredentials = true;

const backend = axios.create({
  baseURL: base_backend_url,
  headers: {
    "Content-Type": "application/json",
  },
});

export { backend };