<template>
    <div>
      <pagetitle
        :title="title"/>
        <widgetsRow :config="layout.widgetRowConfig"/>
    </div>
</template>

<script>

import widgetRowConfig from '@/layout/widgetRowConfig';

export default {
  name: 'home-view',
  data () {
    return {
      app: process.env.APP_NAME,
      title: "Current "+process.env.APP_NAME+" stats",
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
    pagetitle: () => import('@/components/pagetitle'),
    widgetsRow: () => import('@/components/widgetsRow'),
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
</style>