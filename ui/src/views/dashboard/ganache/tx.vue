<template>
    <div>
      <pagetitle/>
      <p>Ganache transactions</p>
    </div>
</template>

<script>

export default {
  name: 'ganache-tx-view',
  data () {
    return {
      app: process.env.APP_NAME,
    }
  },
  methods: {
  },
  created(){
    log("ganache-tx-view::created");
  },
  mounted(){
    log("ganache-tx-view::mounted");
  },
  components: {
    pagetitle: () => import('@/components/pagetitle'),
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
</style>