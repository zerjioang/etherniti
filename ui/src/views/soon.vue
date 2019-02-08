<template>
  <div class="t-root">
    <errorMessage
    class="t-child"
    buttonVisible="false"
    pagetitle="etherniti.org | COOMING SOON"
    code="Cooming Soon"
    :title='apptitle'
    description="Cooming soon available on the web">
      <div slot="image">
        <img :src="require('@/assets/images/ethereum.png')" class="logo">
      </div>
    </errorMessage>
  </div>
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
.t-root {
  display: table;
  height: 100% !important;
  width: 100% !important;
}
.t-child {
  display: table-cell;
  vertical-align: middle;
}
.logo {
  margin: 10px;
  max-width: 50%;
  width: 200px;
}
</style>