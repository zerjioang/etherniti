<template>
    <baseDashboardView>
        <div slot="content">
            <p>Current METHW Platform stats</p>
            <widgetsRow :config="layout.widgetRowConfig"/>
        </div>
    </baseDashboardView>
</template>

<script>

import widgetRowConfig from '@/layout/widgetRowConfig';

export default {
  name: 'index-view',
  data () {
    return {
      title: process.env.UI_TITLE,
      layout: {
        widgetRowConfig: widgetRowConfig,
        version: "0.0.1",
        copyright: "METHW Project",
        years: "2018 - 2019"
      }
    }
  },
  methods: {
  },
  created(){
    log("index-view::created");
  },
  mounted(){
    log("index-view::mounted");
  },
  components: {
    widgetsRow: () => import('@/components/widgetsRow'),
    baseDashboardView: () => import('@/components/baseDashboardView')
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>