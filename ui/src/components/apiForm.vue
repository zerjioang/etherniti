<template>
  <jsonForm
  :form="form"
  :model="model"
  v-on:jsonevent="formEvent">
    <div slot="form-top">
      <dualAlert
      class="top-15"
      :visible="result.visible"
      :valid="result.valid"
      :message="result.message">
      </dualAlert>
    </div>
  </jsonForm>
</template>

<script>

import api from '@/mixins/api';

export default {
  name: 'api-form',
  props: {
    form: {
      type: Object,
      default: () => {}
    },
  },
  mixins: [
    api
  ],
  data () {
    return {
      model: undefined,
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
    formEvent: function(event, eventId){
      event.preventDefault();
      if (eventId=="submit") {
        //submit form data to rest api /v1/eth/verify/:address
        this.api()
        .get(this.endpoint(this.form.api.url + this.model.address))
        .then(response => this.onApiResponse(response))
        .catch(err => this.onApiResponseError(err));
      }
    },
    onApiResponseError(response){
      console.log("error response");
      this.result.visible = true;
      this.result.valid = false;
      if (response && response.response && response.response.data) {
        // error response from REST API
        let resp = response.response.data;
        this.result.message = "Error - "+resp.code+", "+resp.msg+". "+resp.details;
      } else {
        this.result.message = "No response given";
      }
    },
    onApiResponse(response){
      this.result.visible = true;
      this.result.valid = response.data.result;
      if(this.result.valid){
        this.result.message = this.model.address + this.result.messageValid;
      } else {
        this.result.message = this.model.address + this.result.messageInvalid;
      }
    }
  },
  created(){
    log("api-form::created");
    // model deep copy
    log("deep copying form model");
    log(this.form.model);
    this.model = Object.assign({}, this.form.model);
    log("deep copying result");
    log(this.model);
  },
  mounted(){
    log("api-form::mounted");
  },
  components: {
    dualAlert: () => import('@/components/dualAlert'),
    jsonForm: () => import('@/components/jsonForm')
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
.top-15 {
  margin-top: 15px;
}
</style>