<template>
    <div>
      <pagetitle/>
      <p>Ganache blocks</p>
    </div>
</template>

<script>

export default {
  name: 'ganache-blocks-view',
  data () {
    return {
      app: process.env.APP_NAME,
    }
  },
  methods: {
  },
  created(){
    log("ganache-blocks-view::created");
  },
  mounted(){
    log("ganache-blocks-view::mounted");
  },
  components: {
    pagetitle: () => import('@/components/pagetitle'),
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
</style>