import Router from 'vue-router';
import routerNames from '@/layout/routerNames';

//allowed views
const index = () => import('@/views/index');
const license = () => import('@/views/license');
const about = () => import('@/views/about');
const notfound = () => import('@/views/notfound');

const base = process.env.DASHBOARD_BASE_PATH;

export default new Router({
  routes: [
    {
      // home
      path: routerNames.home.path,
      name: routerNames.home.name,
      component: index
    },
    {
      // license
      path: routerNames.license.path,
      name: routerNames.license.name,
      component: license
    },
    {
      // about
      path: routerNames.about.path,
      name: routerNames.about.name,
      component: about
    },
    {
      // not found
      path: routerNames.notfound.path,
      name: routerNames.notfound.name,
      component: notfound
    }
  ],
  mode: process.env.ROUTER_MODE
})
