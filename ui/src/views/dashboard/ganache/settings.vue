<template>
    <div>
      <pagetitle/>
      <p>Ganache settings</p>
    </div>
</template>

<script>

export default {
  name: 'ganache-settings-view',
  data () {
    return {
      app: process.env.APP_NAME,
    }
  },
  methods: {
  },
  created(){
    log("ganache-settings-view::created");
  },
  mounted(){
    log("ganache-settings-view::mounted");
  },
  components: {
    pagetitle: () => import('@/components/pagetitle'),
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
</style>