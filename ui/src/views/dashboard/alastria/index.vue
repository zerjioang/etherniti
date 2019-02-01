<template>
    <div>
      <pagetitle
        title="Alastria Blockchain Ecosystem"
        subtitle="National Blockchain"/>

        <div class="row clearfix">
        </div>
    </div>
</template>

<script>

export default {
  name: 'alastria-index-view',
  data () {
    return {
    }
  },
  methods: {
  },
  created(){
    log("alastria-index-view::created");
  },
  mounted(){
    log("alastria-index-view::mounted");
  },
  components: {
    pagetitle: () => import('@/components/pagetitle'),
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
</style>