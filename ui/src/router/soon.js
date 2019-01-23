import Router from 'vue-router';
import routerNames from '@/layout/routerNames';

//allowed views
const soon = () => import('@/views/soon');
const notfound = () => import('@/views/notfound');

export default new Router({
  routes: [
    {
      path: routerNames.home.path,
      name: routerNames.home.name,
      component: soon
    },
    {
      path: routerNames.notfound.path,
      name: routerNames.notfound.name,
      component: notfound
    }
  ],
  mode: process.env.ROUTER_MODE
})
