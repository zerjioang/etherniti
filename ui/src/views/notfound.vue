<template>
  <errorMessage
  pagetitle="error 404 - methw.org"
  code="404"
  title='Oops.. You just found an error page..'
  description="We are sorry but our service is currently not available, page might have been moved, deleted or abducted by aliens!"></errorMessage>
</template>

<script>

export default {
  name: 'notfound-view',
  data () {
    return {
    }
  },
  methods: {
  },
  created(){
    log("notfound-view::created");
  },
  mounted(){
    log("notfound-view::created");
  },
  components: {
    'errorMessage': () => import('@/components/error')
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
</style>