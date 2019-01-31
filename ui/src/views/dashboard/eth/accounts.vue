<template>
    <div>
      <pagetitle
        title="Create new Ethereum address"
        subtitle="Create new Ethereum and Quorum addresses on the fly"/>

        <div class="row clearfix">
          <apiForm :form="form.addressCheckForm"></apiForm>
        </div>
    </div>
</template>

<script>

import addressCheckForm from '@/layout/forms/addressCheckForm';
export default {
  name: 'eth-accounts-view',
  data () {
    return {
      form: {
        addressCheckForm: addressCheckForm,
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
    log("eth-accounts-view::created");
  },
  mounted(){
    log("eth-accounts-view::mounted");
  },
  components: {
    pagetitle: () => import('@/components/pagetitle'),
    apiForm: () => import('@/components/apiForm'),
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
</style>