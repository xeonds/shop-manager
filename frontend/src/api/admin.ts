import serviceAxios from '../axios';

const customerAPI = {
  BASE_URL: '/admin',
  login: async (username: string, password: string) => serviceAxios.post(`${customerAPI.BASE_URL}/login`, { username, password }).then(response => response.data),
};

export default customerAPI;
