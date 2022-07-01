import axios from 'axios';

const axiosWithAuth = idToken => {
  const authToken = localStorage.getItem("token");
  
  return axios.create({
    baseURL: process.env.REACT_APP_API_URL,
    headers: {
      Authorization: 'Bearer ' + idToken,
    },
  });
};

export default axiosWithAuth;