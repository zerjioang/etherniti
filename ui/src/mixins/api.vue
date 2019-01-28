<template>
</template>

<script type="text/javascript">

//import and register axios
import axios from 'axios';

//configure our target API
const API_ENDPOINT = process.env.API_SCHEME + "://" + process.env.API_DOMAIN + process.env.API_PATH;

export default {
  name: 'api-mixin',
  data () {
    return {
      endpointUrl: undefined
    }
  },
  methods: {
    api(){
      return axios;
    },
    endpoint(path){
      return this.endpointUrl+path;
    }
  },
  created(){
    log("api-mixin::created");
    log("setting default REST API endpoint to: "+API_ENDPOINT);
    this.endpointUrl = API_ENDPOINT;
  },
  mounted(){
    log("api-mixin::mounted");
  },
  components: {
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
</style>