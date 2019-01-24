<!--this file should not be modified -->
<template>
  <div class="four-zero-four err">
    <div class="four-zero-four-container">
        <h1 class="title blue">{{code}}</h1>

        <slot name="image"></slot>

        <h3 class="message text-muted">{{title}}</h3>
        <div class="details text-muted">{{description}}</div>
        
        <slot name="content"></slot>
        
        <div class="button-place">
            <a class="btn btn-default btn-lg waves-effect" v-on:click="onBack">Go back</a>
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
.err {
    padding: 5%;
}
.blue {
  color: #383A3F;
}
.title {
  font-size: 60pt;
}
.message {
  font-size: 25pt;
}
.details {
  font-size: 20pt;
}
</style>