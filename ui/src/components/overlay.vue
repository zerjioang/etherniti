<template>
<div class="overlay"></div>
</template>

<script>
export default {
  name: 'overlay',
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
  watch: {
    visible: function(now, before){
      if(this.visible==true){
        this.show();
      } else {
        this.hide();
      }
    }
  },
  methods: {
    show: function(){
    	log("showing overlay");
    	this.overlay.fadeIn();
    },
    hide: function(){
    	log("removing overlay");
    	this.overlay.fadeOut();
    }
  },
  created(){
    log("overlay::created");
  },
  mounted(){
    log("overlay::mounted");
    this.overlay = $('.overlay');
  },
  components: {
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
</style>