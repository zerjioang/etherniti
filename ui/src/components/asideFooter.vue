<template>
	<!-- Footer -->
	<div class="legal">
	    <div class="copyright">
	        &copy; {{years}} <a href="javascript:void(0);">{{copyright}}</a>.
	    </div>
	    <div class="version">
	        <b>Version: </b> {{version}}
	    </div>
	</div>
</template>

<script>

export default {
  name: 'aside-footer',
  props: {
    version: {
      type: String,
      default: "0.0.1"
    },
    years: {
      type: String,
      default: "2018 - 2019"
    },
    copyright: {
      type: String,
      default: "@zerjioang"
    },
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
    log("aside-footer::created");
  },
  mounted(){
    log("aside-footer::mounted");
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
</style>
