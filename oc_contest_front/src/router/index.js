import { createWebHistory, createRouter } from "vue-router";
import Home from "@/components/Home.vue";
import Self from "@/components/Self.vue";
import Create from "@/components/Create.vue";

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
  },
  {
    path: "/self",
    name: "Self",
    component: Self,
  },
  {
    path: "/create",
    name: "Create",
    component: Create,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;