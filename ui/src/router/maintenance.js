import Router from 'vue-router';
import routerNames from '@/layout/routerNames';

export default new Router({
  routes: [
    {
      path: routerNames.maintenance.path,
      name: routerNames.maintenance.name,
      component: routerNames.maintenance.component
    },
    {
      path: routerNames.notfound.path,
      name: routerNames.notfound.name,
      component: routerNames.notfound.component
    }
  ],
  mode: process.env.ROUTER_MODE
})
