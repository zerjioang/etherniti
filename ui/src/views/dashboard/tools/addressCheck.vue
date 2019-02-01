<template>
    <div>
      <pagetitle
        title="Ethereum Address checker"
        subtitle="Verify whether a given Address is valid or not, check if address is smart contract, etc."/>

        <div class="row clearfix">
          <apiForm :form="form.addressCheckForm"></apiForm>
          <apiForm :form="form.contractAddressCheckForm"></apiForm>
        </div>

        <div class="row clearfix">
          <apiForm :form="form.balanceCheck"></apiForm>
          <apiForm :form="form.balanceCheckAtBlock"></apiForm>
        </div>
        
    </div>
</template>

<script>

import addressCheckForm from '@/layout/forms/addressCheckForm';
import contractAddressCheckForm from '@/layout/forms/contractAddressCheckForm';

import balanceCheckForm from '@/layout/forms/balanceCheckForm';
import balanceAtBlockCheckForm from '@/layout/forms/balanceAtBlockCheckForm';

export default {
  name: 'eth-addr-check-view',
  data () {
    return {
      form: {
        addressCheckForm: addressCheckForm,
        contractAddressCheckForm: contractAddressCheckForm,
        balanceCheck: balanceCheckForm,
        balanceCheckAtBlock: balanceAtBlockCheckForm,
      },
      result: {
        visible: false,
        valid: false,
        messageValid: " is a valid ETH address.",
        messageInvalid: " is an invalid ETH address.",
        message: ""
      }
    }
  },
  methods: {
  },
  created(){
    log("eth-addr-check-view::created");
  },
  mounted(){
    log("eth-addr-check-view::mounted");
  },
  components: {
    pagetitle: () => import('@/components/pagetitle'),
    apiForm: () => import('@/components/apiForm'),
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
.top-15 {
  margin-top: 15px;
}
</style>