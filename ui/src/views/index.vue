<template>
</template>

<script>

import routerNames from '@/layout/routerNames';
import localstorage from '@/mixins/localstorage';

export default {
  name: 'index-view',
  mixins: [
    localstorage
  ],
  data () {
    return {
    }
  },
  methods: {
  },
  created(){
    log("index-view::created");
    if (this.supportsLocalStorage()) {
      log("redirecting to dashboard...");
      this.$router.push({name: routerNames.dashboardHome.name});
    } else {
      // redirect to localstorage error
      log("redirecting to local storage error view...");
      this.$router.push({name: routerNames.localstorage.name});
    }
  },
  mounted(){
    log("index-view::mounted");
  },
  components: {
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
</style>