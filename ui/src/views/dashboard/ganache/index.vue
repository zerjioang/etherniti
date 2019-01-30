<template>
    <div>
      <pagetitle/>
      <p>Ganache index</p>
    </div>
</template>

<script>

export default {
  name: 'ganache-view',
  data () {
    return {
      app: process.env.APP_NAME,
    }
  },
  methods: {
  },
  created(){
    log("ganache-view::created");
  },
  mounted(){
    log("ganache-view::mounted");
  },
  components: {
    pagetitle: () => import('@/components/pagetitle'),
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
</style>