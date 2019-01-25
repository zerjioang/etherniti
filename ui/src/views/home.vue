<template>
    <div>
        <p>Current gaethway stats</p>
        <widgetsRow :config="layout.widgetRowConfig"/>
    </div>
</template>

<script>

import widgetRowConfig from '@/layout/widgetRowConfig';

export default {
  name: 'home-view',
  data () {
    return {
      layout: {
        widgetRowConfig: widgetRowConfig
      }
    }
  },
  methods: {
  },
  created(){
    log("home-view::created");
  },
  mounted(){
    log("home-view::mounted");
  },
  components: {
    widgetsRow: () => import('@/components/widgetsRow'),
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>