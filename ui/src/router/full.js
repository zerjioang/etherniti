import Router from 'vue-router';
import routerNames from '@/layout/routerNames';

//allowed views
const index = () => import('@/views/index');
const notfound = () => import('@/views/notfound');

const base = process.env.DASHBOARD_BASE_PATH;

export default new Router({
  routes: [
    {
      path: routerNames.home.path,
      name: routerNames.home.name,
      component: index
    },
    {
      path: routerNames.notfound.path,
      name: routerNames.notfound.name,
      component: notfound
    }
  ],
  mode: process.env.ROUTER_MODE
})
