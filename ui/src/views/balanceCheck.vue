<template>
    <div>
      <pagetitle
        title="Ethereum Address balance check"
        subtitle="Given a Ethereum Address, check its current balance in real time"/>

        <div class="row clearfix">
          <apiForm :form="form.balanceCheck"></apiForm>
        </div>
    </div>
</template>

<script>

import balanceCheckForm from '@/layout/forms/balanceCheckForm';

export default {
  name: 'eth-addr-balance-view',
  data () {
    return {
      app: process.env.APP_NAME,
      form: {
        balanceCheck: balanceCheckForm
      }
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
    apiForm: () => import('@/components/apiForm'),
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
.centered {
  text-align: center;
}
</style>