<template>
    <div>
        <h2>Ethereum Address Balance Check</h2>
        <p>Check address ETH balance</p>

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
  name: 'eth-addr-balance-view',
  data () {
    return {
      title: process.env.UI_TITLE,
    }
  },
  methods: {
  },
  created(){
    log("eth-addr-balance-view::created");
  },
  mounted(){
    log("eth-addr-balance-view::mounted");
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