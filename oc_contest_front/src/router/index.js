import { createWebHashHistory, createRouter } from "vue-router";
import Home from "@/components/Home.vue";
import Top from "@/components/Top.vue";
import Self from "@/components/Self.vue";
import Create from "@/components/Create.vue";

const routes = [
  {
    path: "/",
    name: "Home",
    meta: {
      title: "OC Contest - Home",
    },
    component: Home,
  },
  {
    path: "/top",
    name: "Top",
    meta: {
      title: "OC Contest - Top",
    },
    component: Top,
  },
  {
    path: "/self",
    name: "Self",
    meta: {
      title: "OC Contest - Own",
    },
    component: Self,
  },
  {
    path: "/create",
    name: "Create",
    meta: {
      title: "OC Contest - Create",
    },
    component: Create,
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

router.afterEach((to) => {
  document.title = to.meta.title || "OC Contest";
});

export default router;