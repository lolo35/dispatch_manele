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
    },
    {
        path: "/closedispatches",
        name: "CloseDispatches",
        component: () => import('../views/CloseDispatchesView.vue')
    },
    {
        path: "/changedispatchtype",
        name: "ChangeDispatchType",
        component: () => import('../views/ChangeDispatchTypeView.vue')
    },
    {
        path: "/updateTargetQty",
        name: "UpdateTargetQty",
        component: () => import('../views/UpdateTargetQtyView.vue')
    }
];

const router = createRouter({
    history: createWebHashHistory(),
    routes,
});

export default router;