<template>
  <errorMessage
  pagetitle="cooming Soon - gaethway.org"
  code="Cooming Soon"
  :title='apptitle'
  description="Cooming soon available on the web"></errorMessage>
</template>

<script>

export default {
  name: 'soon-view',
  data () {
    return {
      apptitle: process.env.APP_TITLE
    }
  },
  methods: {
  },
  created(){
    log("soon-view::created");
  },
  mounted(){
    log("soon-view::mounted");
  },
  components: {
    errorMessage: () => import('@/components/error')
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
</style>