import axios from 'axios';

const productAPI = {
  BASE_URL: '/products',
  getProducts: async () => axios.get(productAPI.BASE_URL).then(response => response.data),
  getProductById: async (id: number) => axios.get(`${productAPI.BASE_URL}/${id}`).then(response => response.data),
  createProduct: async (product: any) => axios.post(productAPI.BASE_URL, product).then(response => response.data),
  updateProduct: async (id: number, product: any) => axios.put(`${productAPI.BASE_URL}/${id}`, product).then(response => response.data),
  deleteProduct: async (id: number) => axios.delete(`${productAPI.BASE_URL}/${id}`).then(response => response.data),
};

export default productAPI;