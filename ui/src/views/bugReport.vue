<template>
    <div>
       <pagetitle
        title="Report an issue"
        subtitle="Have you found a bug? Do you think something can be improved? Let us know!"/>

        <div class="row">
            <div class="col-lg-12 col-md-12 col-sm-12 col-xs-12">
                <div class="card">
                </div>
            </div>
        </div>
    </div>
</template>

<script>

export default {
  name: 'report-bug-view',
  data () {
    return {
      title: process.env.APP_NAME,
    }
  },
  methods: {
  },
  created(){
    log("report-bug-view::created");
  },
  mounted(){
    log("report-bug-view::mounted");
  },
  components: {
    pagetitle: () => import('@/components/pagetitle'),
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
.centered {
  text-align: center;
}
</style>