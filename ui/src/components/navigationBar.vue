<template>
	<nav class="navbar">
		<div class="container-fluid">
		    <div class="navbar-header">
		        <a href="javascript:void(0);" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar-collapse" aria-expanded="false"></a>
		        <a href="javascript:void(0);" class="bars"></a>
		        <router-link class="navbar-brand" :to="{ name: routerNames.dashboardHome.name }">{{title}}</router-link>
		    </div>
		    <div class="collapse navbar-collapse" id="navbar-collapse">
		        <ul class="nav navbar-nav navbar-right">
		            <!-- Call Search -->
		            <li><a href="#" v-on:click="$emit('openSearch')" class="js-search" data-close="true"><i class="material-icons">search</i></a></li>
		            <!-- #END# Tasks -->
		            <li class="pull-right">
		            	<a href="javascript:void(0);" data-close="true" v-on:click="$emit('toggleRightAside')">
		            		<i class="material-icons">more_vert</i>
		            	</a>
		            </li>
		        </ul>
		    </div>
		</div>
		</nav>
</template>

<script>

import routerNames from '@/layout/routerNames';

export default {
  name: 'navigation-bar',
  props: {
  	title: {
  		type: String,
  		default: "NAVBAR"
  	},
  	showSearchBar: {
  		type: Boolean,
  		default: false
  	}
  },
  data () {
    return {
    	routerNames
    }
  },
  methods: {
  	activate: function () {
  		var $body = $('body');
        var $overlay = $('.overlay');

        //Open left sidebar panel
        $('.bars').on('click', function () {
            $body.toggleClass('overlay-open');
            if ($body.hasClass('overlay-open')) { $overlay.fadeIn(); } else { $overlay.fadeOut(); }
        });

        //Close collapse bar on click event
        $('.nav [data-close="true"]').on('click', function () {
            var isVisible = $('.navbar-toggle').is(':visible');
            var $navbarCollapse = $('.navbar-collapse');

            if (isVisible) {
                $navbarCollapse.slideUp(function () {
                    $navbarCollapse.removeClass('in').removeAttr('style');
                });
            }
        });
  	}
  },
  created(){
    log("navigation-bar::created");
  },
  mounted(){
    log("navigation-bar::mounted");
    this.activate();
  },
  components: {
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
</style>