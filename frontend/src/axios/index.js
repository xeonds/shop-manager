import axios from "axios";

const serviceAxios = axios.create({
  baseURL: "/api/",
});
// serviceAxios .interceptors.request.use(
//   (req) => {
//     if (!req.needToken) return req;

//     const token = localStorage.getItem("token");
//     if (!token) return Promise.reject("No token");
//     const data = jwtDecode(token);
//     if (data.exp <= Date.now() / 1000) return Promise.reject("Token expired");
//     req.headers.Authorization = token;

//     return req;
//   },
//   (err) => {
//     return Promise.reject(err);
//   }
// );
// serviceAxios.interceptors.response.use(
//   (req) => {
//     return req;
//   },
//   (err) => {
//     return Promise.reject(err);
//   }
// );

export default serviceAxios;
