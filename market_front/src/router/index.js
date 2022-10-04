import { createWebHistory, createRouter } from "vue-router";
import Home from "@/components/Home.vue";
import Search from "@/components/Search.vue";

const routes = [
  {
    path: "/",
    name: "Home",
    meta: {
      title: "RWBY Adventures Market",
    },
    component: Home,
  },
  {
    path: "/search",
    name: "Search",
    meta: {
      title: "Market - Search",
    },
    component: Search,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.afterEach((to) => {
  document.title = to.meta.title || "RWBY Adventures Market";
});

export default router;