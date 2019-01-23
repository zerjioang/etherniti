<!--this file should not be modified -->
<template>
  <div class="page">
     <div class="page-content">
        <div class="container text-center">
           <div class="display-1 text-muted mb-5"><i class="si si-exclamation"></i>{{code}}</div>
           <h1 class="h2 mb-3">{{title}}</h1>
           <p class="h4 text-muted font-weight-normal mb-7">{{description}}</p>
           <div class="inline">
            <a class="btn btn-primary spaced" href="#" v-on:click="onBack">
              <i class="fe fe-arrow-left mr-2"></i>Go back
            </a>
            <slot name="content"></slot>
           </div>
        </div>
     </div>
  </div>
</template>

<script>


export default {
  name: 'error',
  props: {
    pagetitle: {
      type: String,
      default: "Error page"
    },
    code: {
      type: String,
      default: "503"
    },
    title: {
      type: String,
      default: "Oops.. You just found an error page.."
    },
    description: {
      type: String,
      default: "We are sorry but our service is currently not available"
    }
  },
  mixins: [
  ],
  data () {
    return {
    	uititle: process.env.UI_TITLE,
    	domain: process.env.WEBAPP_DOMAIN
    }
  },
  methods: {
    //goes back on the router. same as history.back()
    onBack: function(e) {
      log("going back in the router");
      this.$router.go(-1);
    }
  },
  created(){
    log("error::created");
  },
  mounted(){
    log("error::mounted");
    document.title = this.pagetitle;
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
.inline {
  display: -webkit-inline-box;
}
.spaced {
  margin: 10px;
}
</style>