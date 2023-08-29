import CreateVue from "@/components/Create.vue";
import List from "@/components/List.vue";
import { createWebHistory, createRouter } from "vue-router";
import { RouteRecordRaw } from "vue-router";
const routes: Array<RouteRecordRaw> =[
    {
        path:'/',
        name :"List",
        component:List
    },
    {
        path:'/add',
        name :"NewBook",
        component:CreateVue
    },
]
const router =createRouter ({
    history: createWebHistory(),
    routes
})
export default router;