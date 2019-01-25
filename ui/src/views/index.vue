<template>
</template>

<script>

import routerNames from '@/layout/routerNames';

export default {
  name: 'index-view',
  data () {
    return {
    }
  },
  methods: {
  },
  created(){
    log("index-view::created");
    log("redirecting to dahboard...");
    this.$router.push({name: routerNames.dashboardHome.name});
  },
  mounted(){
    log("index-view::mounted");
  },
  components: {
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>