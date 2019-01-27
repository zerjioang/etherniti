<template>
    <div>
      <pagetitle
        title="Ethereum address checker"
        subtitle="Verify whether a given adress is valid or not, check if address is smart contract, etc."/>

        <div class="row clearfix">
          <jsonForm
          :config="addressCheckForm"
          v-on:jsonevent="formEvent"></jsonForm>
        </div>
    </div>
</template>

<script>

import addressCheckForm from '@/layout/forms/addressCheckForm';

export default {
  name: 'eth-addr-check-view',
  data () {
    return {
      addressCheckForm: addressCheckForm,
      result: {
        visible: false,
        valid: false,
        messageValid: "is a valid ETH address.",
        messageInvalid: "is an invalid ETH address.",
        message: ""
      }
    }
  },
  methods: {
    formEvent: function(event, eventId){
      event.preventDefault();
      alert("address check");
      alert(eventId);
      log(addressCheckForm);
    },
    show: function(){
      this.result.visible = true;
      if (new Date().getMilliseconds() % 2 == 0) {
        // valid
        this.result.valid = true;
        this.result.message = this.result.messageValid;
      } else {
        //invalid
        this.result.valid = false;
        this.result.message = this.result.messageInvalid;
      }
    },
    reset: function(){
      this.form.address = "";
    }
  },
  created(){
    log("eth-addr-check-view::created");
  },
  mounted(){
    log("eth-addr-check-view::mounted");
  },
  components: {
    pagetitle: () => import('@/components/pagetitle'),
    jsonForm: () => import('@/components/jsonForm')
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
</style>