import Router from 'vue-router';
import routerNames from '@/layout/routerNames';

//allowed views

const dashboardIndex = () => import('@/views/dashboard/home');
const addressCheck = () => import('@/views/dashboard/tools/addressCheck');
const balanceCheck = () => import('@/views/dashboard/tools/balanceCheck');
const privateApi = () => import('@/views/dashboard/home');
const bugReport = () => import('@/views/dashboard/bugReport');
const about = () => import('@/views/dashboard/about');
const license = () => import('@/views/dashboard/license');

const base = process.env.DASHBOARD_BASE_PATH;

export default new Router({
  routes: [
    {
      // index
      path: routerNames.index.path,
      name: routerNames.index.name,
      component: routerNames.index.component
    },
    {
      // dashboard
      path: routerNames.baseDashboardView.path,
      name: routerNames.baseDashboardView.name,
      component: routerNames.baseDashboardView.component,
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
          // index >> dashboard
          path: routerNames.dashboardIndex.path,
          name: routerNames.dashboardIndex.name,
          component: routerNames.dashboardIndex.component
        },
        {
          // dashboard >> address check
          path: routerNames.addressChecker.path,
          name: routerNames.addressChecker.name,
          component: addressCheck
        },
        {
          // dashboard >> ganache accounts
          path: routerNames.ganacheAccounts.path,
          name: routerNames.ganacheAccounts.name,
          component: routerNames.ganacheAccounts.component,
        },
        {
          // dashboard >> ganache transactions
          path: routerNames.ganacheTransactions.path,
          name: routerNames.ganacheTransactions.name,
          component: routerNames.ganacheTransactions.component,
        },
        {
          // dashboard >> ganache blocks
          path: routerNames.ganacheBlocks.path,
          name: routerNames.ganacheBlocks.name,
          component: routerNames.ganacheBlocks.component,
        },
        {
          // dashboard >> ganache settings
          path: routerNames.ganacheSettings.path,
          name: routerNames.ganacheSettings.name,
          component: routerNames.ganacheSettings.component,
        },
        {
          // dashboard >> balance check
          path: routerNames.balanceChecker.path,
          name: routerNames.balanceChecker.name,
          component: balanceCheck
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
      // no local storage
      path: routerNames.localstorage.path,
      name: routerNames.localstorage.name,
      component: routerNames.localstorage.component
    },
    {
      // not found
      path: routerNames.notfound.path,
      name: routerNames.notfound.name,
      component: routerNames.notfound.component
    }
  ],
  scrollBehavior: function (to, from, savedPosition) {
    return {x: 0, y: 0}; // return to top
  },
  mode: process.env.ROUTER_MODE
})
