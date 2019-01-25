<template>
    <div>
        <h2>Report an issue</h2>
        <p>Have you found a bug? Do you think something can be improved? Let us know!</p>

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
      title: process.env.UI_TITLE,
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
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.centered {
  text-align: center;
}
</style>