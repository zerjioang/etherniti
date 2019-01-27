import global from '@/config/global'
import Vue from 'vue';
import App from '@/App';
import selector from '@/router/selector';

import * as OfflinePluginRuntime from 'offline-plugin/runtime';

const prod = process.env.PRODUCTION === true;
if(prod){
	OfflinePluginRuntime.install();
}
Vue.config.productionTip = !prod;

// import vue tour
// info: https://github.com/pulsardev/vue-tour
//import VueTour from 'vue-tour'
//require('vue-tour/dist/vue-tour.css')
//Vue.use(VueTour)

//import and register vue-select
//import vSelect from 'vue-select';
//Vue.component('v-select', vSelect);

//import and register axios
//import axios from 'axios';
//window.axios = axios;

//import select2 from 'select2';
//window.select2 = select2;

//import modal library
import VModal from 'vue-js-modal'
Vue.use(VModal, { dialog: true, dynamic: true, injectModalsContainer: true });

import Router from 'vue-router';
Vue.use(Router);

//configure our router
let router = selector.getRouter();

// load global material js files
// Jquery Core Js
// require('@/assets/plugins/jquery/jquery.min.js');

// Bootstrap Core Js
require('@/assets/plugins/bootstrap/js/bootstrap.js');

// Select Plugin Js
require('@/assets/plugins/bootstrap-select/js/bootstrap-select.js');

// Slimscroll Plugin Js
require('@/assets/plugins/jquery-slimscroll/jquery.slimscroll.js');

// Waves Effect Plugin Js
require('@/assets/plugins/node-waves/waves.js');

// Jquery CountTo Plugin Js
// require('@/assets/plugins/jquery-countto/jquery.countTo.js');

// Morris Plugin Js
// require('@/assets/plugins/raphael/raphael.min.js');
// require('@/assets/plugins/morrisjs/morris.js');

// ChartJs
// require('@/assets/plugins/chartjs/Chart.bundle.js');

// Flot Charts Plugin Js
// require('@/assets/plugins/flot-charts/jquery.flot.js');
// require('@/assets/plugins/flot-charts/jquery.flot.resize.js');
// require('@/assets/plugins/flot-charts/jquery.flot.pie.js');
// require('@/assets/plugins/flot-charts/jquery.flot.categories.js');
// require('@/assets/plugins/flot-charts/jquery.flot.time.js');

// Sparkline Chart Plugin Js
// require('@/assets/plugins/jquery-sparkline/jquery.sparkline.js');

// Custom Js

// require('@/assets/js/pages/index.js');

// Demo Js
// require('@/assets/js/demo.js');

// load global material theme stylesheets

// Bootstrap Core Css
require("@/assets/plugins/bootstrap/css/bootstrap.css");
// Waves Effect Css
require("@/assets/plugins/node-waves/waves.css");
// Animation Css
require("@/assets/plugins/animate-css/animate.css");
// Morris Chart
// require("@/assets/plugins/morrisjs/morris.css");
// Custom Css
require("@/assets/css/style.css");

// AdminBSB Themes. You can choose a theme from css/themes instead of get all themes
require("@/assets/css/themes/theme-gaethway.css");

// load custom made modifications
require("@/assets/css/custom.css");

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  template: '<App/>',
  components: {
  	App
  }
})
