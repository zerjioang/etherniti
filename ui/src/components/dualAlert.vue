<template>
<div v-show="visible" class="col-md-12">
  <div class="alert fade"
	  :class="{
	  'alert-success': valid,
	  'alert-danger': !valid,
	  'in': visible
	  }">
      {{message}}
  </div>
</div>
</template>


<script>

export default {
  name: 'dual-alert',
  props: {
    message: {
      type: String,
      default: "An error occurred while processing your request"
    },
    valid: {
      type: Boolean,
      default: false
    },
    visible: {
      type: Boolean,
      default: false
    }
  },
  mixins: [
  ],
  data () {
    return {
    }
  },
  methods: {
  },
  created(){
    log("dual-alert::created");
  },
  mounted(){
    log("dual-alert::mounted");
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
</style>