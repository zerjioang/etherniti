<template>
    <div>
      <pagetitle
        :title="title"
        subtitle="Following, all used tools for building etherniti are listed with their corresponding licenses."/>
        
        <h3 class="title">Licenses related to UI and this website</h3>
        <licenseInfoCardsRow :config="lists.ui"/>
        
        <h3 class="title">Licenses related to webAPI, tracker and other backend software</h3>
        <licenseInfoCardsRow :config="lists.api"/>
    </div>
</template>

<script>

import licenseList from '@/layout/licenseList';

export default {
  name: 'license-view',
  data () {
    return {
      app: process.env.APP_NAME,
      title: process.env.APP_NAME + " license",
      lists: licenseList
    }
  },
  methods: {
  },
  created(){
    log("license-view::created");
  },
  mounted(){
    log("license-view::mounted");
  },
  components: {
    pagetitle: () => import('@/components/pagetitle'),
    licenseInfoCardsRow: () => import('@/components/licenseInfoCardsRow'),
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
.title {
  padding-bottom: 10px;
}
</style>