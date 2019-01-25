import Router from 'vue-router';
import routerNames from '@/layout/routerNames';

//allowed views
const index = () => import('@/views/index');

const baseDashboard = () => import('@/components/baseDashboardView');

const home = () => import('@/views/home');
const bugReport = () => import('@/views/bugReport');
const about = () => import('@/views/about');
const license = () => import('@/views/license');
const notfound = () => import('@/views/notfound');

const base = process.env.DASHBOARD_BASE_PATH;

export default new Router({
  routes: [
    {
      // index
      path: routerNames.index.path,
      name: routerNames.index.name,
      component: index
    },
    {
      // dashboard
      path: routerNames.dashboardHome.path,
      name: routerNames.dashboardHome.name,
      component: baseDashboard,
      beforeEnter: (to, from, next) => {
        // evaluate before entering to dashboard
        // if the browser supports HTML5 webstorage apis
        const supportsLocalStorageApi = true;
        if (supportsLocalStorageApi) {
          next();
        } else {
          next(
            {
              path: routerNames.notfound.path,
              query: { redirect: to.fullPath }
            }
          )
        }
      },
      children: [
        {
          // dashboard >> index
          path: routerNames.home.path,
          name: routerNames.home.name,
          component: home
        },
        
        {
          // dashboard >> report >> issue
          path: routerNames.bugReport.path,
          name: routerNames.bugReport.name,
          component: bugReport
        },
        {
          // dashboard >> license
          path: routerNames.license.path,
          name: routerNames.license.name,
          component: license
        },
        {
          // dashboard >> about
          path: routerNames.about.path,
          name: routerNames.about.name,
          metadata: { errorView: true },
          component: about
        },
      ]
    },
    {
      // not found
      path: routerNames.notfound.path,
      name: routerNames.notfound.name,
      component: notfound
    },
  ],
  scrollBehavior: function (to, from, savedPosition) {
    return {x: 0, y: 0}; // return to top
  },
  mode: process.env.ROUTER_MODE
})
