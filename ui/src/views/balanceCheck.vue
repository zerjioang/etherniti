<template>
    <div>
      <pagetitle
        title="Ethereum Address balance check"
        subtitle="Given a Ethereum Address, check its current balance in real time"/>

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
      app: process.env.APP_NAME,
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