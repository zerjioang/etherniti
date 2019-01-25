<template>
    <div>
        <h2>gaethway license</h2>
        <p>Following, all used tools for building gaethway are listed with their corresponding licenses.</p>
        <licenseInfoCardsRow/>
    </div>
</template>

<script>
export default {
  name: 'license-view',
  data () {
    return {
      title: process.env.UI_TITLE,
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
    licenseInfoCardsRow: () => import('@/components/licenseInfoCardsRow'),
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>