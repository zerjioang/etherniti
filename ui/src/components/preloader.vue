<template>
	<div class="preloader pl-size-xl" v-if="visible">
	    <div class="spinner-layer pl-indigo">
	        <div class="circle-clipper left">
	            <div class="circle"></div>
	        </div>
	        <div class="circle-clipper right">
	            <div class="circle"></div>
	        </div>
	    </div>
	</div>
</template>


<script>

//import and register axios
import axios from 'axios';

export default {
  name: 'preloader',
  props: {
  	visible: {
  		type: Boolean,
  		default: true
  	}
  },
  data () {
    return {
    }
  },
  methods: {
  },
  created(){
    log("preloader::created");
  },
  mounted(){
    log("preloader::mounted");
  },
  components: {
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
</style>