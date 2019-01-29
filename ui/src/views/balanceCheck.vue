<template>
    <div>
      <pagetitle
        title="Ethereum Address balance"
        subtitle="Given a Ethereum Address, check its current balance."/>

        <div class="row clearfix">
          <apiForm :form="form.balanceCheck"></apiForm>
          <apiForm :form="form.balanceCheckAtBlock"></apiForm>
        </div>
    </div>
</template>

<script>

import balanceCheckForm from '@/layout/forms/balanceCheckForm';
import balanceAtBlockCheckForm from '@/layout/forms/balanceAtBlockCheckForm';

export default {
  name: 'eth-addr-balance-view',
  data () {
    return {
      app: process.env.APP_NAME,
      form: {
        balanceCheck: balanceCheckForm,
        balanceCheckAtBlock: balanceAtBlockCheckForm,
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