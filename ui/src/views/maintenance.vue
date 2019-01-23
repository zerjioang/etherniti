<template>
  <errorMessage
  pagetitle="maintenance - methw.org"
  code="Under maintenance"
  :title='apptitle'
  description="We will back in a flash"></errorMessage>
</template>

<script>

export default {
  name: 'maintenance-view',
  data () {
    return {
      apptitle: process.env.APP_TITLE
    }
  },
  methods: {
  },
  created(){
    log("maintenance-view::created");
  },
  mounted(){
    log("maintenance-view::mounted");
  },
  components: {
    'errorMessage': () => import('@/components/error')
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
</style>