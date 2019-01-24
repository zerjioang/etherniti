<template>
<!-- Page Loader -->
    <div class="page-loader-wrapper">
        <div class="loader">
            <div class="preloader">
                <div class="spinner-layer pl-red">
                    <div class="circle-clipper left">
                        <div class="circle"></div>
                    </div>
                    <div class="circle-clipper right">
                        <div class="circle"></div>
                    </div>
                </div>
            </div>
            <p>{{text}}</p>
        </div>
    </div>
</template>

<script>
export default {
  name: 'loader',
  props: {
  	text: {
  		type: String,
  		default: "Please wait..."
  	},
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
      $('.page-loader-wrapper').fadeIn();
    },
    hide: function(){
      $('.page-loader-wrapper').fadeOut();
    }
  },
  created(){
    log("loader::created");
  },
  mounted(){
    log("loader::mounted");
    setTimeout(this.hide, 50);
  },
  components: {
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
</style>