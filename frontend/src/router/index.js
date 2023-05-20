import { createRouter, createWebHashHistory } from "vue-router";
import Login from "../views/Login.vue";
import Home from "../views/Home.vue";

export default createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: "/",
      redirect: "/login",
    },
    {
      path: "/404",
      name: "NotFound",
      component: () => import("../views/404.vue"),
    },
    {
      path: "/login",
      component: Login,
    },
    {
      path: "/main",
      component: Home,
      children: [
        {
          path: "dashboard",
          component: () => import("../views/home/Dashboard.vue"),
        },
        {
          path: "customer",
          component: () => import("../views/home/Customer.vue"),
        },
        {
          path: "store",
          component: () => import("../views/home/Store.vue"),
        },
      ],
    },
  ],
});
