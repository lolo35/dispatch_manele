import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router';
import HomeView from '../views/HomeView.vue'

const routes:Array<RouteRecordRaw> = [
    {
        path: "/",
        name: "Home",
        component : HomeView,
    },
    {
        path: "/delete",
        name: "Delete",
        component: () => import('../views/DeleteView.vue'),
    },
    {
        path: "/descriptions",
        name: "DispatchDescriptions",
        component: () => import('../views/DispatchDescriptions.vue'),
    }
];

const router = createRouter({
    history: createWebHashHistory(),
    routes,
});

export default router;