<template>
	<!-- Footer -->
	<div class="legal">
	    <div class="copyright">
	        &copy; {{years}} <a href="javascript:void(0);" v-on:click="versionPopup">{{copyright}}</a>.
	    </div>
	    <div class="version">
	        <b>Version: </b> {{version}}
	    </div>
	</div>
</template>

<script>

import aboutDialog from '@/components/dialog/about';
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
  data () {
    return {
    }
  },
  methods: {
    versionPopup: function() {
      this.$modal.show(aboutDialog,
        {
          title: "About Gaethway Project",
          message: 'Gaethway Project is a Multitenant High Performance Ethereum and Quorum compatible WebAPI',
          icon: "code",
          buttonText: "Close"
        }, {
          draggable: false,
          scrollable: true,
          adaptative: true,
          height: "auto"
        },{
          'before-close': (event) => {
            log('modal closed');
          }
        });
    }
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
