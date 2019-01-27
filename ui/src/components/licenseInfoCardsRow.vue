<template>
<div class="row clearfix">
    <div
    v-for="license in config"
    :key="license.id"
    class="col-lg-4 col-md-4 col-sm-6 col-xs-12">
        <licenseInfoCard :license="license"/>
    </div>
</div>
</template>

<script>

export default {
  name: 'license-info-card-row',
  props: {
    config: {
      type: Array,
      default: () => []
    }
  },
  data () {
    return {
    }
  },
  methods: {
  },
  created(){
    log("license-info-card-row::created");
  },
  mounted(){
    log("license-info-card-row::mounted");
  },
  components: {
  	licenseInfoCard: () => import('@/components/licenseInfoCard'),
    baseDashboardView: () => import('@/components/baseDashboardView')
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
</style>