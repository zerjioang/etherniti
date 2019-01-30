import Router from 'vue-router';
import routerNames from '@/layout/routerNames';

export default new Router({
  routes: [
    {
      path: routerNames.soon.path,
      name: routerNames.soon.name,
      component: routerNames.soon.component
    },
    {
      path: routerNames.notfound.path,
      name: routerNames.notfound.name,
      component: routerNames.notfound.component
    }
  ],
  mode: process.env.ROUTER_MODE
})
